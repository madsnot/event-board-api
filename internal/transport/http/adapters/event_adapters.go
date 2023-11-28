package adapters

import (
	"github.com/google/uuid"
	"github.com/madsnot/event-board-api/internal/domain/models"
	"github.com/madsnot/event-board-api/internal/transport/http/dto"
	"time"
)

func AdaptCreateEventRequestToEventBmodel(edto dto.CreateEventRequest) (models.Event, error) {
	id, err := uuid.Parse(edto.AuthorID)
	if err != nil {
		return models.Event{}, err
	}

	start, err := time.Parse(edto.StartDate, "2006-01-02T15:04:05+07:00")
	if err != nil {
		return models.Event{}, err
	}

	var endPtr *time.Time

	if edto.EndDate != "" {
		end, err := time.Parse(edto.StartDate, "2006-01-02T15:04:05+07:00")
		if err != nil {
			return models.Event{}, err
		}

		endPtr = &end
	}

	return models.Event{
		AuthorID:    id,
		Status:      convertStatus(edto.Status),
		Title:       edto.Title,
		Type:        convertType(edto.Type),
		Theme:       edto.Theme,
		Description: edto.Description,
		Genders:     convertGenders(edto.Genders),
		Age:         edto.Age,
		StartDate:   start,
		EndDate:     endPtr,
	}, nil
}

func convertStatus(s int) models.EventStatus {
	switch s {
	case 0:
		return models.EventStatusActive
	default:
		return models.EventStatusClosed
	}
}

func convertType(t int) models.EventType {
	switch t {
	case 0:
		return models.EventTypeOnline
	default:
		return models.EventTypeOffline
	}
}

func convertGenders(g dto.EventGender) models.EventGender {
	return models.EventGender{
		Man:   g.Man,
		Woman: g.Woman,
	}
}
