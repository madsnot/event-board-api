package usecase

import (
	"context"
	"github.com/madsnot/event-board-api/internal/domain/models"
	"github.com/madsnot/event-board-api/internal/repository"
)

type UserUsecase struct {
	rep repository.UserRepository
}

func NewUserUsecase(rep repository.UserRepository) *UserUsecase {
	return &UserUsecase{
		rep: rep,
	}
}

func (uu UserUsecase) GetUserByEmail(ctx context.Context, email string) (models.User, error) {
	user, err := uu.rep.GetUserByEmail(ctx, email)
	if err != nil {
		return models.User{}, ErrUserNotFound.Wrap(err)
	}

	return user, nil
}
