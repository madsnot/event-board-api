package v1

import (
	"github.com/madsnot/event-board-api/internal/transport/http/dto"
	"net/http"

	"github.com/madsnot/event-board-api/internal/transport/http/adapters"
)

func (h *handler) GetEventList(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	req, err := parseRequestToGetEventsRequest(r)
	if err != nil {
		writeError(w, err)
		return
	}

	filters, err := adapters.AdaptGetEventsRequestToFilters(req)
	if err != nil {
		writeError(w, err)
		return
	}

	list, err := h.usecase.Event.GetList(ctx, filters)
	if err != nil {
		writeError(w, err)
		return
	}

	writeOK(w, adapters.AdaptEventsToGetEventsResponse(list))
}

func (h *handler) GetEvent(w http.ResponseWriter, r *http.Request) {

}

func (h *handler) CreateEvent(w http.ResponseWriter, r *http.Request) {
	req, err := parseRequestToCreateEventRequest(r)
	if err != nil {
		writeError(w, err)
		return
	}

	ctx := r.Context()

	model, err := adapters.AdaptCreateEventRequestToEventBmodel(req)
	if err != nil {
		writeError(w, err)
		return
	}

	if err = h.usecase.Event.CreateEvent(ctx, model); err != nil {
		writeError(w, err)
		return
	}

	writeOK(w, dto.CreateEventResponse{})
}

func (h *handler) UpdateEvent(w http.ResponseWriter, r *http.Request) {

}

func (h *handler) CloseEvent(w http.ResponseWriter, r *http.Request) {

}
