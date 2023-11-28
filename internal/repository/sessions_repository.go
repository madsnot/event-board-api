package repository

import (
	"context"
	"github.com/madsnot/event-board-api/internal/domain/models"
	"github.com/madsnot/event-board-api/pkg/database"
	"github.com/rs/zerolog"
)

type SessionRepository struct {
	db  database.INoSqlDb
	log zerolog.Logger
}

func NewSessionRepository(db database.INoSqlDb, log zerolog.Logger) *SessionRepository {
	return &SessionRepository{
		db:  db,
		log: log,
	}
}

func (sr SessionRepository) CreateSession(ctx context.Context, session models.Session) (models.Session, error) {
	//dto := adaptBmodelToDTO(session)

	return models.Session{}, nil
}
