package postgres

import (
	"encoding/json"
	"github.com/google/uuid"
	"github.com/jackc/pgtype"
	"github.com/madsnot/event-board-api/internal/domain/models"
	dto2 "github.com/madsnot/event-board-api/internal/repository/postgres/dto"
	"time"
)

func adaptSessionBmodelToDTO(session models.Session) dto2.SessionDatabaseDTO {
	return dto2.SessionDatabaseDTO{
		UserID:   session.UserID.String(),
		TimeZone: session.TimeZone,
	}
}

func adaptUserBmodelToDTO(user models.User) dto2.UserDatabaseDTO {
	return dto2.UserDatabaseDTO{}
}

func adaptUserDTOToBmodel(user dto2.UserDatabaseDTO) models.User {
	var bDate *time.Time

	if user.BirthdayDate.Status == pgtype.Null {
		bDate = &user.BirthdayDate.Time
	}

	return models.User{
		ID:           user.ID.Get().(uuid.UUID),
		Username:     user.Username,
		Email:        user.Email,
		Password:     user.Password,
		Avatar:       user.Avatar,
		Firstname:    user.Firstname,
		Lastname:     user.Lastname,
		Middlename:   user.Middlename,
		Gender:       convertDTOGenderToBmodel(user.Gender),
		BirthdayDate: bDate,
		CreatedAt:    user.CreatedAt,
		UpdatedAt:    user.UpdatedAt,
	}
}

func adaptEventBmodelToDTO(event models.Event) (dto2.EventDTO, error) {
	eventDTO := dto2.EventDTO{
		Status:      convertBmodelStatusToDTO(event.Status),
		Title:       event.Title,
		Theme:       event.Theme,
		Type:        convertBmodelTypeToDTO(event.Type),
		Description: event.Description,
		Age:         event.Age,
		StartDate:   event.StartDate,
		EndDate:     event.EndDate,
	}

	if err := eventDTO.AuthorID.Set(event.AuthorID); err != nil {
		return dto2.EventDTO{}, err
	}

	if err := eventDTO.Genders.Set(event.Genders); err != nil {
		return dto2.EventDTO{}, err
	}

	return eventDTO, nil
}

func adaptEventDTOToBmodel(edto dto2.EventDTO) (models.Event, error) {
	genders, err := convertDTOGendersToBmodel(edto.Genders)
	if err != nil {
		return models.Event{}, err
	}

	return models.Event{
		ID:          edto.ID.Bytes,
		Status:      convertDTOStatusToBmodel(edto.Status),
		AuthorID:    edto.AuthorID.Bytes,
		Title:       edto.Title,
		Type:        convertDTOTypeToBmodel(edto.Type),
		Theme:       edto.Theme,
		Description: edto.Description,
		Genders:     genders,
		Age:         edto.Age,
		StartDate:   edto.StartDate,
		EndDate:     edto.EndDate,
		CreatedAt:   edto.CreatedAt,
		UpdatedAt:   edto.UpdatedAt,
	}, nil
}

func convertBmodelTypeToDTO(etype models.EventType) dto2.EventType {
	switch etype {
	case models.EventTypeOnline:
		return dto2.EventTypeOnline
	case models.EventTypeOffline:
		return dto2.EventTypeOffline
	default:
		return dto2.EventTypeEmpty
	}
}

func convertBmodelStatusToDTO(status models.EventStatus) dto2.EventStatus {
	switch status {
	case models.EventStatusNew:
		return dto2.EventStatusNew
	case models.EventStatusActive:
		return dto2.EventStatusActive
	case models.EventStatusClosed:
		return dto2.EventStatusClosed
	default:
		return dto2.EventStatusEmpty
	}
}

func convertDTOTypeToBmodel(etype dto2.EventType) models.EventType {
	switch etype {
	case dto2.EventTypeOnline:
		return models.EventTypeOnline
	case dto2.EventTypeOffline:
		return models.EventTypeOffline
	default:
		return models.EventTypeEmpty
	}
}

func convertDTOStatusToBmodel(status dto2.EventStatus) models.EventStatus {
	switch status {
	case dto2.EventStatusNew:
		return models.EventStatusNew
	case dto2.EventStatusActive:
		return models.EventStatusActive
	case dto2.EventStatusClosed:
		return models.EventStatusClosed
	default:
		return models.EventStatusEmpty
	}
}

func convertDTOGendersToBmodel(jsonb pgtype.JSONB) (models.EventGender, error) {
	var genders models.EventGender

	if err := json.Unmarshal(jsonb.Bytes, &genders); err != nil {
		return models.EventGender{}, err
	}

	return genders, nil
}

func convertDTOGenderToBmodel(g int) models.GenderType {
	switch g {
	case 0:
		return models.Woman
	default:
		return models.Man
	}
}
