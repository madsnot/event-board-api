package dto

type EventGender struct {
	Man   bool `json:"man"`
	Woman bool `json:"women"`
}

type CreateEventRequest struct {
	Status      int         `json:"status"`
	AuthorID    string      `json:"authorId"`
	Title       string      `json:"title"`
	Type        int         `json:"type"`
	Theme       string      `json:"theme"`
	Description string      `json:"description"`
	Genders     EventGender `json:"genders"`
	Age         int         `json:"age"`
	StartDate   string      `json:"startDate"`
	EndDate     string      `json:"endDate"`
}
