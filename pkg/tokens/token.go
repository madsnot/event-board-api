package tokens

import (
	"crypto/rand"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type TokenManager interface {
	CreateToken(userId string) (*Token, error)
}

type Token struct {
	AccessToken  string
	RefreshToken string
}

type TokenInfo struct {
	accessTokenTTL  time.Duration
	refreshTokenTTL time.Duration
	signingKey      string
}

func NewTokenInfo(accessTokenTTL time.Duration, refreshTokenTTL time.Duration, signingKey string) *TokenInfo {
	return &TokenInfo{
		accessTokenTTL:  accessTokenTTL,
		refreshTokenTTL: refreshTokenTTL,
		signingKey:      signingKey,
	}
}

func (token *TokenInfo) NewAccessToken(tokenTemp string) (signedAccessToken string, err error) {
	newAccessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		ExpiresAt: time.Now().Add(token.accessTokenTTL).Unix(),
		Subject:   tokenTemp,
	})
	signedAccessToken, err = newAccessToken.SignedString([]byte(token.signingKey))
	if err != nil {
		return "", err
	}
	return signedAccessToken, nil
}

func (token *TokenInfo) NewRefreshToken() (refreshToken string, err error) {
	newRefreshToken := make([]byte, 15)
	_, err = rand.Read(newRefreshToken)
	if err != nil {
		return "", err
	}
	refreshToken = string(newRefreshToken)
	return refreshToken, nil
}
