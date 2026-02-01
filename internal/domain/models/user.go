package models

import (
	"time"

	"github.com/google/uuid"
)

type UserGenderType int

const (
	UserGenderTypeUnknown UserGenderType = iota
	UserGenderTypeWoman
	UserGenderTypeMan
)

type User struct {
	ID           uuid.UUID
	Username     string
	Email        string
	Password     string
	Avatar       string
	Firstname    string
	Lastname     string
	Middlename   string
	Gender       UserGenderType
	BirthdayDate *time.Time
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
