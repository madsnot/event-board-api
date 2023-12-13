package dto

type EventGenderDTO struct {
	Man   bool `json:"man"`
	Woman bool `json:"woman"`
}

type EventDTO struct {
	ID          string         `json:"id"`
	Status      int            `json:"status"`
	AuthorID    string         `json:"authorId"`
	Title       string         `json:"title"`
	Type        int            `json:"type"`
	Theme       string         `json:"theme"`
	Description string         `json:"description"`
	Genders     EventGenderDTO `json:"genders"`
	Age         int            `json:"age"`
	StartDate   string         `json:"startDate"`
	EndDate     string         `json:"endDate"`
	CreateAt    string         `json:"createAt"`
	UpdateAt    string         `json:"updateAt"`
	ClosedAt    string         `json:"closedAt"`
}

type GetEventsRequest struct {
	Query     string         `json:"query"`
	Statuses  []int          `json:"statuses"`
	AuthorIDs []string       `json:"authorIds"`
	Type      int            `json:"type"`
	Themes    []string       `json:"themes"`
	Genders   EventGenderDTO `json:"genders"`
	Age       int            `json:"age"`
	Older     bool           `json:"older"`
	Younger   bool           `json:"younger"`
	StartDate string         `json:"startDate"`
	EndDate   string         `json:"endDate"`
	CreateAt  string         `json:"createAt"`
}

type GetEventsResponse struct {
	Events []EventDTO
}

type CreateEventRequest struct {
	Status      int            `json:"status"`
	AuthorID    string         `json:"authorId"`
	Title       string         `json:"title"`
	Type        int            `json:"type"`
	Theme       string         `json:"theme"`
	Description string         `json:"description"`
	Genders     EventGenderDTO `json:"genders"`
	Age         int            `json:"age"`
	StartDate   string         `json:"startDate"`
	EndDate     string         `json:"endDate"`
}

type CreateEventResponse struct {
}
