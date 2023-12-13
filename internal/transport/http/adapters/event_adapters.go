package adapters

import (
	"github.com/google/uuid"
	"github.com/madsnot/event-board-api/internal/domain/models"
	"github.com/madsnot/event-board-api/internal/transport/http/dto"
	"time"
)

func AdaptGetEventsRequestToFilters(gedto dto.GetEventsRequest) (models.EventFilters, error) {
	var (
		startPtr, endPtr, createPtr *time.Time
		ids                         []uuid.UUID
		statuses                    []models.EventStatus
		themes                      []string
	)

	for _, val := range gedto.AuthorIDs {
		id, err := uuid.Parse(val)
		if err != nil {
			return models.EventFilters{}, err
		}

		ids = append(ids, id)
	}

	for _, val := range gedto.Statuses {
		statuses = append(statuses, convertDTOToStatus(val))
	}

	if len(gedto.Themes) > 0 {
		themes = gedto.Themes
	}

	if gedto.Genders.Man == false && gedto.Genders.Woman == false {
		gedto.Genders.Man = true
		gedto.Genders.Woman = true
	}

	if gedto.StartDate != "" {
		start, err := time.Parse(time.RFC3339, gedto.StartDate)
		if err != nil {
			return models.EventFilters{}, err
		}

		startPtr = &start
	}

	if gedto.EndDate != "" {
		end, err := time.Parse(time.RFC3339, gedto.EndDate)
		if err != nil {
			return models.EventFilters{}, err
		}

		endPtr = &end
	}

	if gedto.CreateAt != "" {
		create, err := time.Parse(time.RFC3339, gedto.CreateAt)
		if err != nil {
			return models.EventFilters{}, err
		}

		createPtr = &create
	}

	return models.EventFilters{
		Query:     gedto.Query,
		Statuses:  statuses,
		AuthorIDs: ids,
		Type:      convertDTOToType(gedto.Type),
		Themes:    themes,
		Genders:   convertDTOToGenders(gedto.Genders),
		Age:       gedto.Age,
		Older:     gedto.Older,
		Younger:   gedto.Younger,
		StartDate: startPtr,
		EndDate:   endPtr,
		CreatedAt: createPtr,
	}, nil
}

func AdaptCreateEventRequestToEventBmodel(edto dto.CreateEventRequest) (models.Event, error) {
	id, err := uuid.Parse(edto.AuthorID)
	if err != nil {
		return models.Event{}, err
	}

	start, err := time.Parse(time.RFC3339, edto.StartDate)
	if err != nil {
		return models.Event{}, err
	}

	var endPtr *time.Time

	if edto.EndDate != "" {
		end, err := time.Parse(time.RFC3339, edto.EndDate)
		if err != nil {
			return models.Event{}, err
		}

		endPtr = &end
	}

	return models.Event{
		AuthorID:    id,
		Status:      convertDTOToStatus(edto.Status),
		Title:       edto.Title,
		Type:        convertDTOToType(edto.Type),
		Theme:       edto.Theme,
		Description: edto.Description,
		Genders:     convertDTOToGenders(edto.Genders),
		Age:         edto.Age,
		StartDate:   start,
		EndDate:     endPtr,
	}, nil
}

func AdaptEventToEventDTO(event models.Event) dto.EventDTO {
	var end string

	if event.EndDate != nil {
		end = event.EndDate.String()
	}

	return dto.EventDTO{
		ID:          event.ID.String(),
		Status:      convertStatusToDTO(event.Status),
		AuthorID:    event.AuthorID.String(),
		Title:       event.Title,
		Type:        convertTypeToDTO(event.Type),
		Theme:       event.Theme,
		Description: event.Description,
		Genders:     convertGendersToDTO(event.Genders),
		Age:         event.Age,
		StartDate:   event.StartDate.String(),
		EndDate:     end,
		CreateAt:    event.CreatedAt.String(),
		UpdateAt:    event.UpdatedAt.String(),
	}
}

func AdaptEventsToGetEventsResponse(list []models.Event) dto.GetEventsResponse {
	events := make([]dto.EventDTO, 0)

	for _, e := range list {
		events = append(events, AdaptEventToEventDTO(e))
	}

	return dto.GetEventsResponse{
		Events: events,
	}
}

func convertDTOToStatus(s int) models.EventStatus {
	switch s {
	case 1:
		return models.EventStatusNew
	case 2:
		return models.EventStatusActive
	case 3:
		return models.EventStatusClosed
	default:
		return models.EventStatusEmpty
	}
}

func convertDTOToType(t int) models.EventType {
	switch t {
	case 1:
		return models.EventTypeOnline
	case 2:
		return models.EventTypeOffline
	default:
		return models.EventTypeEmpty
	}
}

func convertDTOToGenders(g dto.EventGenderDTO) models.EventGender {
	return models.EventGender{
		Man:   g.Man,
		Woman: g.Woman,
	}
}

func convertStatusToDTO(s models.EventStatus) int {
	switch s {
	case models.EventStatusNew:
		return 1
	case models.EventStatusActive:
		return 2
	case models.EventStatusClosed:
		return 3
	default:
		return 0
	}
}

func convertTypeToDTO(t models.EventType) int {
	switch t {
	case models.EventTypeOnline:
		return 1
	case models.EventTypeOffline:
		return 2
	default:
		return 0
	}
}

func convertGendersToDTO(g models.EventGender) dto.EventGenderDTO {
	return dto.EventGenderDTO{
		Man:   g.Man,
		Woman: g.Woman,
	}
}
