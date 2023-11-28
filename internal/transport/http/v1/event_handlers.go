package v1

import (
	"net/http"

	"github.com/madsnot/event-board-api/internal/transport/http/adapters"
)

func (h *handler) GetEventList(w http.ResponseWriter, r *http.Request) {

}

func (h *handler) GetEvent(w http.ResponseWriter, r *http.Request) {

}

func (h *handler) CreateEvent(w http.ResponseWriter, r *http.Request) {
	req, err := parseRequestToCreateEventRequest(r)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	ctx := r.Context()

	model, err := adapters.AdaptCreateEventRequestToEventBmodel(req)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	if err = h.usecase.Event.CreateEvent(ctx, model); err != nil {

	}
}

func (h *handler) UpdateEvent(w http.ResponseWriter, r *http.Request) {

}

func (h *handler) CloseEvent(w http.ResponseWriter, r *http.Request) {

}
