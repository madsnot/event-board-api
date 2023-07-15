package usecase

import (
	repo "github.com/madsnot/event-board-api/internal/repository"
	"github.com/madsnot/event-board-api/pkg/hash"
	"github.com/madsnot/event-board-api/pkg/token"
)

type AuthUsecase struct {
	hasher      *hash.Hasher
	tokenizer   *token.Tokenizer
	userRepo    repo.UserRepositoryInterface
	sessionRepo repo.SessionRepositoryInterface
}

func NewAuthUsecase(hasher *hash.Hasher, tokenizer *token.Tokenizer,
	userRepo repo.UserRepositoryInterface, sessionRepo repo.SessionRepositoryInterface) AuthUsecase {
	return AuthUsecase{
		hasher:      hasher,
		tokenizer:   tokenizer,
		userRepo:    userRepo,
		sessionRepo: sessionRepo,
	}
}
