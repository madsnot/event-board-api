package users

import (
	"example/event-board/server/pkg/email"
	"example/event-board/server/pkg/hash"
	"example/event-board/server/pkg/tokens"
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
