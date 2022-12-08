package users

import (
	"example/event-board/pkg/hash"

	"example/event-board/pkg/tokens"
)

type UserService struct {
	hasher    *hash.Hasher
	tokenInfo *tokens.TokenInfo
}

func NewUserService(hasher *hash.Hasher, tokenInfo *tokens.TokenInfo) *UserService {
	return &UserService{
		hasher:    hasher,
		tokenInfo: tokenInfo,
	}
}
