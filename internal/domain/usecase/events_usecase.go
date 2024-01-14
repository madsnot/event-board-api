package usecase

import (
	"context"
	"fmt"
	"github.com/madsnot/event-board-api/internal/domain/models"
	"github.com/madsnot/event-board-api/internal/repository/opensearch"
	"github.com/madsnot/event-board-api/internal/repository/postgres"
	"github.com/madsnot/event-board-api/internal/repository/rabbit"
)

type EventUsecase struct {
	rep postgres.EventRepository
	os  opensearch.Client
	r   rabbit.Client
}

func NewEventUsecase(rep postgres.EventRepository, os opensearch.Client, r rabbit.Client) *EventUsecase {
	return &EventUsecase{
		rep: rep,
		os:  os,
		r:   r,
	}
}

func (eu EventUsecase) GetList(ctx context.Context, filters models.EventFilters) ([]models.Event, error) {
	var err error

	filters.EventIDs, err = eu.os.Search(ctx, filters.Query)
	if err != nil {
		return nil, ErrDocsNotFound.Wrap(err)
	}

	list, err := eu.rep.GetList(ctx, filters)
	if err != nil {
		return nil, ErrInvalidToGetEvents.Wrap(err)
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
		return ErrInvalidToCreateEvent.Wrap(err)
	}

	if err = eu.os.Index(ctx, event); err != nil {
		return ErrInvalidToCreateIndex.Wrap(err)
	}

	body := []byte(fmt.Sprintf("New event %s was created!", event.Title))

	msg := eu.r.CreateMsg(body)

	if err = eu.r.Publish(ctx, rabbit.CreateRoutingKey, msg); err != nil {
		return ErrInvalidToPublishMsg.Wrap(err)
	}

	return nil
}
