package models

import (
	"time"

	"github.com/google/uuid"
)

type EventFilters struct {
	Query     string
	EventIDs  []uuid.UUID
	Statuses  []EventStatus
	AuthorIDs []uuid.UUID
	Types     []EventType
	Themes    []EventTheme
	Genders   []EventGender
	Age       []EventAge
	StartDate *time.Time
	EndDate   *time.Time
	CreatedAt *time.Time
}
