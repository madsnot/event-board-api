package hash

import (
	"crypto/sha1"
	"fmt"

	"github.com/madsnot/event-board-api/internal/config"
)

type PasswordHasher interface {
	Hash(password string) (string, error)
}

type Hasher struct {
	salt string
}

func NewHasher(cfg config.HashConfig) *Hasher {
	return &Hasher{salt: cfg.HashSalt}
}

func (hasher *Hasher) Hash(password string) (string, error) {
	hash := sha1.New()
	_, err := hash.Write([]byte(password))
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%x", hash.Sum([]byte(hasher.salt))), nil
}
