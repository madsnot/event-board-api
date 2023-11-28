package v1

import (
	"encoding/json"
	"net/http"

	"github.com/madsnot/event-board-api/internal/transport/http/dto"
)

func parseRequestToSignInRequestDTO(r *http.Request) (dto.SignInRequest, error) {
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

func parseRequestToCreateEventRequest(r *http.Request) (dto.CreateEventRequest, error) {
	var (
		body   []byte
		reqDTO dto.CreateEventRequest
	)

	_, err := r.Body.Read(body)
	if err != nil {
		return dto.CreateEventRequest{}, err
	}

	err = json.Unmarshal(body, &reqDTO)
	if err != nil {
		return dto.CreateEventRequest{}, err
	}

	return reqDTO, nil
}
