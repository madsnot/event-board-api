package dto

type UserDTO struct {
	ID           int    `db:"id"`
	Name         string `db:"user_name"`
	Sex          string `db:"sex"`
	BirthdayDate string `db:"birthday_date"`
	Email        string `db:"email"`
}
