package main

import (
	"context"
	"crypto/tls"
	"errors"
	"github.com/gorilla/mux"
	"github.com/madsnot/event-board-api/internal/config"
	"github.com/madsnot/event-board-api/internal/domain/usecase"
	osRep "github.com/madsnot/event-board-api/internal/repository/opensearch"
	"github.com/madsnot/event-board-api/internal/repository/postgres"
	"github.com/madsnot/event-board-api/internal/repository/rabbit"
	httpServer "github.com/madsnot/event-board-api/internal/transport/http/v1"
	"github.com/madsnot/event-board-api/pkg/database"
	"github.com/madsnot/event-board-api/pkg/hash"
	"github.com/madsnot/event-board-api/pkg/token"
	"github.com/madsnot/event-board-api/tern"
	"github.com/opensearch-project/opensearch-go/v2"
	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/rs/zerolog"
	"net"
	"net/http"
	"os"
)

type Server struct {
	httpSrv   *http.Server
	router    *mux.Router
	listener  net.Listener
	cfg       config.Config
	db        database.ISqlDb
	client    database.INoSqlDb
	tokenizer *token.Tokenizer
	hasher    *hash.Hasher
	logger    zerolog.Logger
}

func NewServer() *Server {
	cfg, err := config.LoadConfig()
	if err != nil {
		return &Server{}
	}

	router := mux.NewRouter()

	httpSrv := http.Server{
		Addr:         cfg.Host + ":" + cfg.Port,
		Handler:      router,
		ReadTimeout:  cfg.ReadTimeout,
		WriteTimeout: cfg.WriteTimeout,
	}

	db := database.NewSqlDbClient(cfg.PostgresCfg.DSN)

	tokenizer, err := token.NewTokenizer(cfg.TokenCfg)
	if err != nil {
		return &Server{}
	}

	hasher := hash.NewHasher(cfg.HashCfg)

	return &Server{
		httpSrv:   &httpSrv,
		router:    router,
		cfg:       cfg,
		db:        db,
		tokenizer: tokenizer,
		hasher:    hasher,
	}
}

func (srv *Server) Run(ctx context.Context) error {
	var err error

	if srv.cfg.MigrationsCfg.AutoRun {
		err = tern.RunMigrations(ctx, srv.cfg)
		if err != nil {
			srv.logger.Error().Err(err).Msg("failed to run migrations")
			return err
		}
	}

	srv.logger = zerolog.New(os.Stderr).With().Timestamp().Logger()

	if srv.db != nil {
		if err = srv.db.Open(ctx); err != nil {
			srv.logger.Error().Err(err).Msg("failed to open database conn")
			return err
		}
	}

	srv.listener, err = net.Listen("tcp", srv.httpSrv.Addr)
	if err != nil {
		srv.logger.Error().Err(err).Msg("failed to listen addr")
		return err
	}

	err = srv.initRouters(ctx)
	if err != nil {
		srv.logger.Error().Err(err).Msg("failed init routers")
		return err
	}

	go func() {
		srv.logger.Info().Msg("starting http server on :8082")

		err = srv.httpSrv.Serve(srv.listener)
		if err != nil {
			if errors.Is(err, http.ErrServerClosed) {
				srv.logger.Error().Err(err).Msg("http server closed")
				return
			}

			srv.logger.Fatal().Err(err).Msg("http server error")
		}

		return
	}()

	return nil
}

func (srv *Server) Close() {
	if srv.httpSrv != nil {
		if err := srv.httpSrv.Close(); err != nil {
			srv.logger.Error().Err(err).Msg("failed to close http server")
		}
	}

	if srv.db != nil {
		if err := srv.db.Close(); err != nil {
			srv.logger.Error().Err(err).Msg("failed to close database")
		}
	}
}

func (srv *Server) initRouters(ctx context.Context) error {
	clientOs, err := opensearch.NewClient(opensearch.Config{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
		Addresses: []string{
			srv.cfg.OpensearchConfig.Host,
		},
		Username: srv.cfg.OpensearchConfig.Username,
		Password: srv.cfg.OpensearchConfig.Password,
	})
	if err != nil {
		return err
	}

	osClient := osRep.NewOpensearchClient(srv.cfg.OpensearchConfig, clientOs)

	conn, err := amqp.Dial(srv.cfg.RabbitConfig.Host)
	if err != nil {
		return err
	}

	channel, err := conn.Channel()
	if err != nil {
		return err
	}

	rabbitClient := rabbit.NewClient(conn, channel, srv.cfg.RabbitConfig.Queue, srv.cfg.RabbitConfig.ExchangeName)

	consumer, err := rabbitClient.CreateConsumer(srv.cfg.RabbitConfig.Queue, srv.logger)
	if err != nil {
		return err
	}

	go consumer.SendNotificationToClient(ctx)

	sessionRep := postgres.NewSessionRepository(srv.db, srv.logger)
	userRep := postgres.NewUserRepository(srv.db, srv.logger)
	eventRep := postgres.NewEventRepository(srv.db, srv.logger)

	if err = osClient.InitIndex(ctx, eventRep); err != nil {
		return err
	}

	authUC := usecase.NewAuthUsecase(srv.hasher, srv.tokenizer, userRep, sessionRep)
	userUC := usecase.NewUserUsecase(userRep)
	eventUC := usecase.NewEventUsecase(eventRep, osClient, rabbitClient)

	uc := usecase.NewUsecase(authUC, userUC, eventUC)

	handler := httpServer.NewHandler(uc, srv.cfg)
	handler.Register(srv.router)

	return nil
}
