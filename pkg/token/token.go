package token

import (
	"crypto/rand"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/madsnot/event-board-api/internal/config"
)

type Tokenizer struct {
	accessTokenTTL  time.Duration
	refreshTokenTTL time.Duration
	signingKey      string
}

func NewTokenizer(cfg config.TokenConfig) *Tokenizer {
	return &Tokenizer{
		accessTokenTTL:  cfg.AccessTokenTTL,
		refreshTokenTTL: cfg.RefreshTokenTTL,
		signingKey:      cfg.SigningKey,
	}
}

func (token *Tokenizer) NewAccessToken(tokenTemp string) (signedAccessToken string, err error) {
	newAccessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		ExpiresAt: time.Now().Add(token.accessTokenTTL).Unix(),
		Issuer:    tokenTemp,
	})

	signedAccessToken, err = newAccessToken.SignedString([]byte(token.signingKey))
	if err != nil {
		return "", err
	}

	return signedAccessToken, nil
}

func (token *Tokenizer) NewRefreshToken() (refreshToken string, err error) {
	newRefreshToken := make([]byte, 15)

	_, err = rand.Read(newRefreshToken)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%x", newRefreshToken), nil
}
