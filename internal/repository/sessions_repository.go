package repository

import (
	"context"
	"github.com/madsnot/event-board-api/pkg/database"
	"golang.org/x/exp/slog"
)

type SessionRepository struct {
	db  database.DBInterface
	log slog.Logger
}

func NewSessionRepository(db database.DBInterface, log slog.Logger) *SessionRepository {
	return &SessionRepository{
		db:  db,
		log: log,
	}
}

func (sr SessionRepository) CreateSession(ctx context.Context, userId int, refreshToken string) error {
	return nil
}
