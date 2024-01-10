package models

import (
	"github.com/cristalhq/jwt/v4"
	"github.com/google/uuid"
)

type Claims struct {
	UserID uuid.UUID
	jwt.RegisteredClaims
}
