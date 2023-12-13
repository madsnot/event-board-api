package postgres

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/jackc/pgtype"
	"github.com/madsnot/event-board-api/internal/domain/models"
	"github.com/madsnot/event-board-api/internal/repository"
	"github.com/madsnot/event-board-api/internal/repository/postgres/dto"
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

func (er EventRepository) GetList(ctx context.Context, filters models.EventFilters) ([]models.Event, error) {
	var (
		event dto.EventDTO
		args  []interface{}
	)

	sql := `SELECT id, status, author_id, title, type, theme, description, genders, age, start_date, end_date, created_at 
			FROM events
			WHERE genders = $1`

	var genders pgtype.JSONB

	if err := genders.Set(filters.Genders); err != nil {
		return nil, err
	}

	args = append(args, genders)

	i := 1
	if filters.Younger {
		i++
		sql += fmt.Sprintf(` AND age <= $%d`, i)
		args = append(args, filters.Age)
	} else {
		i++
		sql += fmt.Sprintf(` AND age >= $%d`, i)
		args = append(args, filters.Age)
	}

	if filters.EventIDs != nil {
		i++
		sql += fmt.Sprintf(` AND id = ANY($%d)`, i)
		args = append(args, filters.EventIDs)
	}

	if filters.AuthorIDs != nil {
		i++
		sql += fmt.Sprintf(` AND author_id = ANY($%d)`, i)
		args = append(args, filters.AuthorIDs)
	}

	if filters.Statuses != nil {
		i++
		sql += fmt.Sprintf(` AND status = ANY($%d)`, i)
		args = append(args, filters.Statuses)
	}

	if filters.Themes != nil {
		sql += fmt.Sprintf(` AND theme = ANY($%d)`, i)
		args = append(args, filters.Themes)
	}

	if filters.Type != models.EventTypeEmpty {
		sql += fmt.Sprintf(` AND type = $%d`, i)
		args = append(args, filters.Type)
	}

	if filters.StartDate != nil {
		sql += fmt.Sprintf(` AND start_date = $%d`, i)
		args = append(args, filters.StartDate)
	}

	if filters.EndDate != nil {
		sql += fmt.Sprintf(` AND end_date = $%d`, i)
		args = append(args, filters.EndDate)
	}

	if filters.CreatedAt != nil {
		sql += fmt.Sprintf(` AND created_at = $%d`, i)
		args = append(args, filters.CreatedAt)
	}

	rows, err := er.db.Query(ctx, sql, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	list := make([]models.Event, 0)

	for rows.Next() {
		err = rows.Scan(
			&event.ID,
			&event.Status,
			&event.AuthorID,
			&event.Title,
			&event.Type,
			&event.Theme,
			&event.Description,
			&event.Genders,
			&event.Age,
			&event.StartDate,
			&event.EndDate,
			&event.CreatedAt,
		)
		if err != nil {
			return nil, err
		}

		bevent, err := repository.adaptEventDTOToBmodel(event)
		if err != nil {
			return nil, err
		}

		list = append(list, bevent)
	}

	return list, rows.Err()
}

func (er EventRepository) CreateEvent(ctx context.Context, event models.Event) (uuid.UUID, error) {
	dto, err := repository.adaptEventBmodelToDTO(event)
	if err != nil {
		return uuid.Nil, err
	}

	tx, err := er.db.BeginTx(ctx)
	if err != nil {
		return uuid.Nil, err
	}
	defer tx.Rollback(ctx)

	var id pgtype.UUID

	if err = tx.QueryRow(
		ctx,
		`INSERT INTO events(status, author_id, title, type, theme, description, genders, age, start_date, end_date, created_at, updated_at)  VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, now(), now()) RETURNING id`,
		dto.Status,
		dto.AuthorID,
		dto.Title,
		dto.Type,
		dto.Theme,
		dto.Description,
		dto.Genders,
		dto.Age,
		dto.StartDate,
		dto.EndDate,
	).Scan(&id); err != nil {
		return uuid.Nil, err
	}

	return id.Bytes, tx.Commit(ctx)
}
