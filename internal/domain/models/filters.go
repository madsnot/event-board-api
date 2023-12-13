package models

import (
	"github.com/google/uuid"
	"time"
)

type EventFilters struct {
	Query     string
	EventIDs  []uuid.UUID
	Statuses  []EventStatus
	AuthorIDs []uuid.UUID
	Type      EventType
	Themes    []string
	Genders   EventGender
	Age       int
	Older     bool
	Younger   bool
	StartDate *time.Time
	EndDate   *time.Time
	CreatedAt *time.Time
}
