package models

import "github.com/gofrs/uuid"

type Session struct {
	ID           uuid.UUID
	UserID       uuid.UUID
	RefreshToken string
}
