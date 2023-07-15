package main

import (
	"context"
	"github.com/madsnot/event-board-api/internal/config"
	"github.com/madsnot/event-board-api/pkg/database"
	"github.com/madsnot/event-board-api/pkg/hash"
	"github.com/madsnot/event-board-api/pkg/token"
	"golang.org/x/sync/errgroup"
)

type Server struct {
	cfg       config.Config
	db        database.DBInterface
	tokenizer *token.Tokenizer
	hasher    *hash.Hasher
}

func NewServer() *Server {
	cfg, err := config.LoadConfig()
	if err != nil {
		return &Server{}
	}

	db := database.NewDB(cfg.DataBaseCfg.DSN)

	tokenizer := token.NewTokenizer(cfg.TokenCfg)

	hasher := hash.NewHasher(cfg.HashCfg)

	return &Server{
		cfg:       cfg,
		db:        db,
		tokenizer: tokenizer,
		hasher:    hasher,
	}
}

func (srv Server) Run(ctx context.Context) error {
	if srv.db != nil {
		if err := srv.db.Open(ctx); err != nil {
			return err
		}
	}

	srv.initServices(ctx)

	g := new(errgroup.Group)

	//g.Go(func() error {
	//	return srv.grpcServer.Start()
	//})

	return g.Wait()
}

func (srv Server) Close(ctx context.Context) {
	_ = srv.db.Close()
}

func (srv Server) initServices(ctx context.Context) {

}
