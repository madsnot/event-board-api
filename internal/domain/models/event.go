package models

import (
	"time"

	"github.com/google/uuid"
)

type EventType int

const (
	EventTypeUnknown EventType = iota
	EventTypeEmpty
	EventTypeOnline
	EventTypeOffline
)

type EventGender int

const (
	EventGenderUnknown EventGender = iota
	EventGenderWoman
	EventGenderMan
)

type EventStatus int

const (
	EventStatusUnknown EventStatus = iota
	EventStatusEmpty
	EventStatusNew
	EventStatusActive
	EventStatusCancled
	EventStatusClosed
)

type EventTheme int

const (
	EventThemeUnknown EventTheme = iota
	EventThemeSport
	EventThemePCGames
	EventThemeTableGames
	EventThemeArt
	EventThemeTV
	EventThemeCulture
	EventThemeHoliday
	EventThemeOther
)

type EventAge int

const (
	EventAgeUnknown EventAge = iota
	EventAgeSix
	EventAgeTwelve
	EventAgeSixteen
	EventAgeEighteen
	EventAgeTwentyOne
)

type Event struct {
	ID          uuid.UUID
	Status      EventStatus
	AuthorID    uuid.UUID
	Title       string
	Type        EventType
	Theme       EventTheme
	Description string
	Genders     []EventGender
	Age         EventAge
	StartDate   time.Time
	EndDate     time.Time
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
