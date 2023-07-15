package repository

import (
	"context"
)

type UserRepositoryInterface interface {
	CreateUser()
	GetUserByEmail()
}

type SessionRepositoryInterface interface {
	CreateSession(ctx context.Context, userId int, refreshToken string) error
}
