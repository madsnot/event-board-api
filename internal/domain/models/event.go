package models

import (
	"github.com/google/uuid"
	"time"
)

type (
	EventType   int
	EventGender struct {
		Man   bool
		Woman bool
	}
	EventStatus int
)

const (
	EventTypeEmpty EventType = iota
	EventTypeOnline
	EventTypeOffline
)

const (
	EventStatusEmpty EventStatus = iota
	EventStatusNew
	EventStatusActive
	EventStatusClosed
)

type Event struct {
	ID          uuid.UUID
	Status      EventStatus
	AuthorID    uuid.UUID
	Title       string
	Type        EventType
	Theme       string
	Description string
	Genders     EventGender
	Age         int
	StartDate   time.Time
	EndDate     *time.Time
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
