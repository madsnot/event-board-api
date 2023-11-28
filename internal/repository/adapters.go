package repository

import (
	"github.com/google/uuid"
	"github.com/jackc/pgtype"
	"github.com/madsnot/event-board-api/internal/domain/models"
	"github.com/madsnot/event-board-api/internal/repository/dto"
	"time"
)

func adaptSessionBmodelToDTO(session models.Session) dto.SessionDatabaseDTO {
	return dto.SessionDatabaseDTO{
		UserID:   session.UserID.String(),
		TimeZone: session.TimeZone,
	}
}

func adaptUserBmodelToDTO(user models.User) dto.UserDatabaseDTO {
	return dto.UserDatabaseDTO{}
}

func adaptUserDTOToBmodel(user dto.UserDatabaseDTO) models.User {
	var bDate *time.Time

	if user.BirthdayDate.Status == pgtype.Null {
		bDate = &user.BirthdayDate.Time
	}

	return models.User{
		ID:           user.ID.Get().(uuid.UUID),
		Name:         user.Name,
		Gender:       user.Gender,
		BirthdayDate: bDate,
		Email:        user.Email,
	}
}
func adaptEventBmodelToDTO(event models.Event) (dto.EventDTO, error) {
	eventDTO := dto.EventDTO{
		Title:       event.Title,
		Theme:       event.Theme,
		Type:        convertEventBmodelTypeToDTO(event.Type),
		Description: event.Description,
		Age:         event.Age,
	}

	if err := eventDTO.AuthorID.Set(event.AuthorID); err != nil {
		return dto.EventDTO{}, err
	}

	if err := eventDTO.Genders.Set(event.Genders); err != nil {
		return dto.EventDTO{}, err
	}

	return eventDTO, nil
}

func convertEventBmodelTypeToDTO(etype models.EventType) dto.EventType {
	switch etype {
	case models.EventTypeOnline:
		return dto.EventTypeOnline
	default:
		return dto.EventTypeOffline
	}
}
