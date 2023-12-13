package dto

import (
	"github.com/jackc/pgtype"
)

type UserDatabaseDTO struct {
	ID           pgtype.UUID `db:"id"`
	Name         string      `db:"user_name"`
	Gender       string      `db:"gender"`
	BirthdayDate pgtype.Date `db:"birthday_date"`
	Email        string      `db:"email"`
}
