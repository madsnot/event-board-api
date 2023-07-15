package repository

import (
	"context"
	"github.com/madsnot/event-board-api/internal/repository/dto"
	"github.com/madsnot/event-board-api/pkg/database"
	"golang.org/x/exp/slog"
)

type UserRepository struct {
	db  database.DBInterface
	log slog.Logger
}

func NewUserRepository(db database.DBInterface, log slog.Logger) UserRepository {
	return UserRepository{
		db:  db,
		log: log,
	}
}

func (ur UserRepository) CreateUser(ctx context.Context, user dto.UserDTO) error {
	return nil
}

func (ur UserRepository) GetUserByEmail(ctx context.Context, email string) (user dto.UserDTO, err error) {
	return dto.UserDTO{}, err
}
