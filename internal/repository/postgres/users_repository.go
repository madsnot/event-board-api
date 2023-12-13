package postgres

import (
	"context"
	"github.com/gofrs/uuid"
	"github.com/madsnot/event-board-api/internal/domain/models"
	"github.com/madsnot/event-board-api/internal/repository"
	"github.com/madsnot/event-board-api/internal/repository/postgres/dto"
	"github.com/madsnot/event-board-api/pkg/database"
	"github.com/rs/zerolog"
)

type UserRepository struct {
	db  database.ISqlDb
	log zerolog.Logger
}

func NewUserRepository(db database.ISqlDb, log zerolog.Logger) UserRepository {
	return UserRepository{
		db:  db,
		log: log,
	}
}

func (ur UserRepository) CreateUser(ctx context.Context, user models.User) (uuid.UUID, error) {
	return uuid.Nil, nil
}

func (ur UserRepository) GetUserByEmail(ctx context.Context, email string) (user models.User, err error) {
	var userDTO dto.UserDatabaseDTO

	tx, err := ur.db.BeginTx(ctx)
	if err != nil {
		return models.User{}, err
	}

	row := tx.QueryRow(ctx, `SELECT * FROM users WHERE email = $1`, email)

	if err = row.Scan(&userDTO); err != nil {

	}

	return repository.adaptUserDTOToBmodel(userDTO), err
}
