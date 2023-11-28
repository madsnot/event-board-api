package usecase

import (
	"context"
	"github.com/madsnot/event-board-api/internal/domain/models"
	"github.com/madsnot/event-board-api/internal/repository"
)

type EventUsecase struct {
	rep repository.EventRepository
}

func NewEventUsecase(rep repository.EventRepository) *EventUsecase {
	return &EventUsecase{
		rep: rep,
	}
}

func (eu EventUsecase) CreateEvent(ctx context.Context, event models.Event) error {
	if err := eu.rep.CreateEvent(ctx, event); err != nil {
		return err
	}

	return nil
}
