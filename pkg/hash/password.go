package hash

import (
	"golang.org/x/crypto/bcrypt"
)

type PasswordHasher interface {
	Hash(password string) (string, error)
}

type Hasher struct {
	addition string
}

func NewHasher(addition string) *Hasher {
	return &Hasher{addition: addition}
}

func (hasher *Hasher) Hash(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		return "", err
	}
	return string(hash) + hasher.addition, nil
}
