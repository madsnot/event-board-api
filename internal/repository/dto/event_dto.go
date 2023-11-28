package dto

import (
	"github.com/jackc/pgtype"
	"time"
)

type EventType int

const (
	EventTypeOnline EventType = iota
	EventTypeOffline
)

type EventDTO struct {
	ID          pgtype.UUID
	AuthorID    pgtype.UUID
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
