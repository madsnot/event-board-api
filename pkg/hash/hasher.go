package hash

import (
	"crypto/sha1"

	"github.com/madsnot/event-board-api/internal/config"
)

type Hasher struct {
	salt string
}

func NewHasher(cfg config.HashConfig) *Hasher {
	return &Hasher{salt: cfg.HashSalt}
}

func (hasher *Hasher) Hash(str string) (string, error) {
	hash := sha1.New()

	_, err := hash.Write([]byte(str))
	if err != nil {
		return "", err
	}

	return string(hash.Sum([]byte(hasher.salt))), nil
}
