package models

type UsersList struct {
	ID           int     `json:"id"`
	Email        string  `json:"email"`
	Password     string  `json:"pass"`
	Name         string  `json:"name"`
	Surname      string  `json:"surname"`
	StudentInfo  Student `json:"sInfo"`
	Sex          string  `json:"sex"`
	BirthdayDate string  `json:"bDate"`
}
