package usecase

import (
	"context"
	"github.com/madsnot/event-board-api/internal/domain/models"
	"github.com/madsnot/event-board-api/internal/repository/opensearch"
	"github.com/madsnot/event-board-api/internal/repository/postgres"
)

type EventUsecase struct {
	rep postgres.EventRepository
	os  opensearch.Client
}

func NewEventUsecase(rep postgres.EventRepository, os opensearch.Client) *EventUsecase {
	return &EventUsecase{
		rep: rep,
		os:  os,
	}
}

func (eu EventUsecase) GetList(ctx context.Context, filters models.EventFilters) ([]models.Event, error) {
	var err error

	filters.EventIDs, err = eu.os.Search(ctx, filters.Query)
	if err != nil {
		return nil, nil
	}

	list, err := eu.rep.GetList(ctx, filters)
	if err != nil {
		return nil, nil
	}

	return list, nil
}

func (eu EventUsecase) CreateEvent(ctx context.Context, event models.Event) error {
	var err error

	if err = validEvent(event); err != nil {
		return ErrInvalidEvent.Wrap(err)
	}

	event.ID, err = eu.rep.CreateEvent(ctx, event)
	if err != nil {
		return err
	}

	if err = eu.os.Index(ctx, event); err != nil {
		return err
	}

	return nil
}
