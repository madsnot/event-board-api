package models

type Event struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Date   string `json:"date"`
	Status string `json:"status"`
}
