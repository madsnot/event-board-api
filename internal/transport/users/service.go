package users

import (
	"github.com/madsnot/event-board-api/pkg/email"
	"github.com/madsnot/event-board-api/pkg/hash"
	"github.com/madsnot/event-board-api/pkg/tokens"
)

type UserService struct {
	hasher    *hash.Hasher
	tokenInfo *tokens.TokenInfo
	email     *email.Email
}

func NewUserService(hasher *hash.Hasher, tokenInfo *tokens.TokenInfo, email *email.Email) *UserService {
	return &UserService{
		hasher:    hasher,
		tokenInfo: tokenInfo,
		email:     email,
	}
}
