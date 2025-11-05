package v1

import (
	"encoding/json"
	"github.com/madsnot/event-board-api/internal/transport/http/dto"
	"io"
	"net/http"
)

func parseRequestToSignInRequestDTO(r *http.Request) (dto.SignInRequest, error) {
	if err := r.ParseForm(); err != nil {
		return dto.SignInRequest{}, err
	}

	return dto.SignInRequest{
		Email:    r.PostForm.Get("email"),
		Password: r.PostForm.Get("password"),
	}, nil
}

func parseSignInResponseToResponse(r *http.Request) (dto.SignInRequest, error) {
	var (
		body   []byte
		reqDTO dto.SignInRequest
	)

	_, err := r.Body.Read(body)
	if err != nil {
		return dto.SignInRequest{}, err
	}

	err = json.Unmarshal(body, &reqDTO)
	if err != nil {
		return dto.SignInRequest{}, err
	}

	return reqDTO, nil
}

func parseRequestToGetEventsRequest(r *http.Request) (dto.GetEventsRequest, error) {
	var (
		req  dto.GetEventsRequest
		body []byte
	)

	body, err := io.ReadAll(r.Body)
	if err != nil {
		return dto.GetEventsRequest{}, err
	}

	if err := json.Unmarshal(body, &req); err != nil {
		return dto.GetEventsRequest{}, err
	}

	return req, nil
}

func parseRequestToCreateEventRequest(r *http.Request) (dto.CreateEventRequest, error) {
	var req dto.CreateEventRequest

	body, err := io.ReadAll(r.Body)
	if err != nil {
		return dto.CreateEventRequest{}, err
	}

	if err := json.Unmarshal(body, &req); err != nil {
		return dto.CreateEventRequest{}, err
	}

	return req, nil
}
