package main

import (
	"context"
	"errors"
	"github.com/gorilla/mux"
	httpServer "github.com/madsnot/event-board-api/internal/transport/http/v1"
	"github.com/madsnot/event-board-api/tern"
	"github.com/rs/zerolog"
	"net"
	"net/http"
	"os"

	"github.com/madsnot/event-board-api/internal/config"
	"github.com/madsnot/event-board-api/internal/domain/usecase"
	"github.com/madsnot/event-board-api/internal/repository"
	"github.com/madsnot/event-board-api/pkg/database"
	"github.com/madsnot/event-board-api/pkg/hash"
	"github.com/madsnot/event-board-api/pkg/token"
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

	tokenizer := token.NewTokenizer(cfg.TokenCfg)

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

	srv.initRouters()

	go func() {
		srv.logger.Info().Msg("starting http server on :8080")

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

func (srv *Server) initRouters() {
	sessionRep := repository.NewSessionRepository(srv.db, srv.logger)
	userRep := repository.NewUserRepository(srv.db, srv.logger)
	eventRep := repository.NewEventRepository(srv.db, srv.logger)

	authUC := usecase.NewAuthUsecase(srv.hasher, srv.tokenizer, userRep, sessionRep)
	userUC := usecase.NewUserUsecase(userRep)
	eventUC := usecase.NewEventUsecase(eventRep)

	uc := usecase.NewUsecase(authUC, userUC, eventUC)

	handler := httpServer.NewHandler(uc, srv.cfg)
	handler.Register(srv.router)
}
