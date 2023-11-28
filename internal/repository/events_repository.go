package repository

import (
	"context"
	"github.com/madsnot/event-board-api/internal/domain/models"
	"github.com/madsnot/event-board-api/pkg/database"
	"github.com/rs/zerolog"
)

type EventRepository struct {
	db  database.ISqlDb
	log zerolog.Logger
}

func NewEventRepository(db database.ISqlDb, log zerolog.Logger) EventRepository {
	return EventRepository{
		db:  db,
		log: log,
	}
}

func (er EventRepository) CreateEvent(ctx context.Context, event models.Event) error {
	dto, err := adaptEventBmodelToDTO(event)
	if err != nil {
		return err
	}

	tx, err := er.db.BeginTx(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	_, err = tx.Query(
		ctx,
		`INSERT INTO events(author_id, title, type, theme, description, genders, age, created_at, updated_at)  VALUES ($1, $2, $3, $4, $5, $6, $7, now(), now())`,
		dto.AuthorID,
		dto.Title,
		dto.Type,
		dto.Theme,
		dto.Description,
		dto.Genders,
		dto.Age,
	)
	if err != nil {
		return err
	}

	return tx.Commit(ctx)
}
