package dto

import (
	"github.com/jackc/pgtype"
	"time"
)

type UserDatabaseDTO struct {
	ID           pgtype.UUID `db:"id"`
	Username     string      `db:"username"`
	Email        string      `db:"email"`
	Password     string      `db:"password"`
	Avatar       string      `db:"avatar_url"`
	Firstname    string      `db:"firstname"`
	Lastname     string      `db:"lastname"`
	Middlename   string      `db:"middlename"`
	Gender       int         `db:"gender"`
	BirthdayDate pgtype.Date `db:"birthday_date"`
	CreatedAt    time.Time   `db:"created_at"`
	UpdatedAt    time.Time   `db:"updated_at"`
}
