package usecase

import (
	"context"
	"github.com/cristalhq/jwt/v4"
	"github.com/google/uuid"
	"github.com/madsnot/event-board-api/internal/domain/models"
	repo "github.com/madsnot/event-board-api/internal/repository/postgres"
	"github.com/madsnot/event-board-api/pkg/hash"
	"github.com/madsnot/event-board-api/pkg/token"
	"time"
)

type AuthUsecase struct {
	hasher     *hash.Hasher
	tokenizer  *token.Tokenizer
	userRep    repo.IUserRepository
	sessionRep repo.ISessionRepository
}

func NewAuthUsecase(hasher *hash.Hasher, tokenizer *token.Tokenizer,
	userRepo repo.IUserRepository, sessionRepo repo.ISessionRepository) *AuthUsecase {
	return &AuthUsecase{
		hasher:     hasher,
		tokenizer:  tokenizer,
		userRep:    userRepo,
		sessionRep: sessionRepo,
	}
}

func (ac AuthUsecase) CreateSession(ctx context.Context, userID uuid.UUID) (models.Session, error) {
	var err error

	session := models.Session{UserID: userID}

	session.RefreshToken, err = ac.tokenizer.NewRefreshToken()
	if err != nil {
		return models.Session{}, err
	}

	session, err = ac.sessionRep.CreateSession(ctx, session)
	if err != nil {
		return models.Session{}, err
	}

	now := time.Now().UTC()

	claims := models.Claims{
		UserID: session.UserID,
	}

	claims.ExpiresAt = jwt.NewNumericDate(now.Add(ac.tokenizer.GetAccessTokenTTL()))
	claims.IssuedAt = jwt.NewNumericDate(now)

	session.AccessToken, err = ac.tokenizer.NewAccessToken(claims)

	return session, nil
}
