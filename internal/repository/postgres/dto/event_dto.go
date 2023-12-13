package dto

import (
	"github.com/jackc/pgtype"
	"time"
)

type (
	EventType   int
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

type EventDTO struct {
	ID          pgtype.UUID
	AuthorID    pgtype.UUID
	Status      EventStatus
	Title       string
	Type        EventType
	Theme       string
	Description string
	Genders     pgtype.JSONB
	Age         int
	StartDate   time.Time
	EndDate     pgtype.Time
	CreatedAt   time.Time
	UpdatedAt   time.Time
	ClosedAt    pgtype.Time
}
