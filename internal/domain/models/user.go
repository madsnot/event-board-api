package models

type User struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	Sex          string `json:"sex"`
	BirthdayDate string `json:"bDate"`
	Email        string `json:"email"`
}
