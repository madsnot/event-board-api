package models

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
)

type Claims struct {
	UserID uuid.UUID
	jwt.StandardClaims
}
