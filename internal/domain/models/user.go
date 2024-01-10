package models

import (
	"github.com/google/uuid"
	"time"
)

type GenderType int

const (
	Woman GenderType = iota
	Man
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
	Gender       GenderType
	BirthdayDate *time.Time
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
