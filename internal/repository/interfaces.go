package repository

import (
	"context"
	"github.com/gofrs/uuid"
	"github.com/madsnot/event-board-api/internal/domain/models"
)

type IUserRepository interface {
	CreateUser(ctx context.Context, user models.User) (uuid.UUID, error)
}

type ISessionRepository interface {
	CreateSession(ctx context.Context, session models.Session) (models.Session, error)
}

type IEventRepository interface {
	CreateEvent(ctx context.Context, event models.Event) error
}
