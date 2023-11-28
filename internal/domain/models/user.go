package models

import (
	"github.com/google/uuid"
	"time"
)

type User struct {
	ID           uuid.UUID
	Name         string
	Gender       string
	BirthdayDate *time.Time
	Email        string
	Password     string
}
