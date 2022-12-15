package hash

import (
	"crypto/sha1"
	"fmt"
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
	hash := sha1.New()
	_, err := hash.Write([]byte(password))
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%x", hash.Sum([]byte(hasher.addition))), nil
}
