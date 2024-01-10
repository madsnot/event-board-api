package postgres

import (
	"context"
	"github.com/gofrs/uuid"
	"github.com/madsnot/event-board-api/internal/domain/models"
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

	row := ur.db.QueryRow(ctx, `SELECT id, username, email, password, avatar_url, firstname, lastname, middlename, gender, birthdate, created_at, updated_at FROM users WHERE email = $1`, email)

	if err = row.Scan(
		&userDTO.ID,
		&userDTO.Username,
		&userDTO.Email,
		&userDTO.Password,
		&userDTO.Avatar,
		&userDTO.Firstname,
		&userDTO.Lastname,
		&userDTO.Middlename,
		&userDTO.Gender,
		&userDTO.BirthdayDate,
		&userDTO.CreatedAt,
		&userDTO.UpdatedAt); err != nil {
		return models.User{}, err
	}

	return adaptUserDTOToBmodel(userDTO), err
}
