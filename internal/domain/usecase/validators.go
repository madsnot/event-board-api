package usecase

import (
	"github.com/madsnot/event-board-api/internal/domain/models"
	//"time"
)

func validEvent(event models.Event) error {
	if event.EndDate != nil && event.StartDate.After(*event.EndDate) {
		return ErrInvalidStartDate
	}

	return nil
}
