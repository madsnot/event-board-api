package usecase

import (
	"github.com/madsnot/event-board-api/pkg/email"
	"github.com/madsnot/event-board-api/pkg/hash"
	"github.com/madsnot/event-board-api/pkg/tokens"
)

type UserUsecase struct {
	hasher    *hash.Hasher
	tokenInfo *tokens.Tokenizer
	email     *email.Email
}

func NewUserUsecase(hasher *hash.Hasher, tokenInfo *tokens.Tokenizer, email *email.Email) *UserUsecase {
	return &UserUsecase{
		hasher:    hasher,
		tokenInfo: tokenInfo,
		email:     email,
	}
}
