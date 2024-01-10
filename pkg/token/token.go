package token

import (
	"crypto/rand"
	"fmt"
	"time"

	"github.com/cristalhq/jwt/v4"
	"github.com/madsnot/event-board-api/internal/config"
)

type Tokenizer struct {
	accessTokenTTL  time.Duration
	refreshTokenTTL time.Duration
	signingKey      string
	builder         *jwt.Builder
	verifier        jwt.Verifier
}

func NewTokenizer(cfg config.TokenConfig) (*Tokenizer, error) {
	signer, err := jwt.NewSignerHS(jwt.HS512, []byte(cfg.SigningKey))
	if err != nil {
		return nil, err
	}

	builder := jwt.NewBuilder(signer)

	return &Tokenizer{
		accessTokenTTL:  cfg.AccessTokenTTL,
		refreshTokenTTL: cfg.RefreshTokenTTL,
		signingKey:      cfg.SigningKey,
		builder:         builder,
		verifier:        signer,
	}, nil
}

func (t *Tokenizer) GetAccessTokenTTL() time.Duration {
	return t.accessTokenTTL
}

func (t *Tokenizer) NewAccessToken(claims any) (signedAccessToken string, err error) {
	newAccessToken, err := t.builder.Build(claims)
	if err != nil {
		return "", err
	}

	return newAccessToken.String(), nil
}

func (t *Tokenizer) NewRefreshToken() (refreshToken string, err error) {
	newRefreshToken := make([]byte, 15)

	_, err = rand.Read(newRefreshToken)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%x", newRefreshToken), nil
}

func (t *Tokenizer) UnmarshalJWT(token string, claims any) error {
	return jwt.ParseClaims([]byte(token), t.verifier, claims)
}
