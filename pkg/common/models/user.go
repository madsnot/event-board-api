package models

type UsersList struct {
	ID           int    `json:"id"`
	Login        string `json:"login"`
	Password     string `json:"pass"`
	Email        string `json:"email"`
	Name         string `json:"name"`
	Surname      string `json:"surname"`
	Sex          string `json:"sex"`
	BirthdayDate int    `json:"bDate"`
}
