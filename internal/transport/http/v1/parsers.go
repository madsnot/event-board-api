package v1

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/madsnot/event-board-api/internal/transport/http/dto"
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
	if err := r.ParseMultipartForm(8192); err != nil {
		return dto.GetEventsRequest{}, err
	}

	var statuses []int

	err := json.Unmarshal([]byte(r.PostForm.Get("statuses")), &statuses)
	if err != nil {
		return dto.GetEventsRequest{}, err
	}

	var ids, themes []string

	err = json.Unmarshal([]byte(r.PostForm.Get("authorIds")), &ids)
	if err != nil {
		return dto.GetEventsRequest{}, err
	}

	err = json.Unmarshal([]byte(r.PostForm.Get("themes")), &themes)
	if err != nil {
		return dto.GetEventsRequest{}, err
	}

	etype, err := strconv.Atoi(r.PostForm.Get("type"))
	if err != nil {
		return dto.GetEventsRequest{}, err
	}

	age, err := strconv.Atoi(r.PostForm.Get("age"))
	if err != nil {
		return dto.GetEventsRequest{}, err
	}

	var genders dto.EventGenderDTO

	data := []byte(r.PostForm.Get("genders"))

	if err = json.Unmarshal(data, &genders); err != nil {
		return dto.GetEventsRequest{}, err
	}

	older, err := strconv.ParseBool(r.PostForm.Get("older"))
	if err != nil {
		return dto.GetEventsRequest{}, err
	}

	younger, err := strconv.ParseBool(r.PostForm.Get("younger"))
	if err != nil {
		return dto.GetEventsRequest{}, err
	}

	return dto.GetEventsRequest{
		Query:     r.PostForm.Get("query"),
		Statuses:  statuses,
		AuthorIDs: ids,
		Type:      etype,
		Themes:    themes,
		Genders:   genders,
		Age:       age,
		Older:     older,
		Younger:   younger,
		StartDate: r.PostForm.Get("startDate"),
		EndDate:   r.PostForm.Get("endDate"),
		CreateAt:  r.PostForm.Get("createAt"),
	}, nil
}

func parseRequestToCreateEventRequest(r *http.Request) (dto.CreateEventRequest, error) {
	if err := r.ParseMultipartForm(8192); err != nil {
		return dto.CreateEventRequest{}, err
	}

	status, err := strconv.Atoi(r.PostForm.Get("status"))
	if err != nil {
		return dto.CreateEventRequest{}, err
	}

	etype, err := strconv.Atoi(r.PostForm.Get("type"))
	if err != nil {
		return dto.CreateEventRequest{}, err
	}

	age, err := strconv.Atoi(r.PostForm.Get("age"))
	if err != nil {
		return dto.CreateEventRequest{}, err
	}

	var genders dto.EventGenderDTO

	data := []byte(r.PostForm.Get("genders"))

	if err = json.Unmarshal(data, &genders); err != nil {
		return dto.CreateEventRequest{}, err
	}

	return dto.CreateEventRequest{
		Status:      status,
		AuthorID:    r.PostForm.Get("authorId"),
		Title:       r.PostForm.Get("title"),
		Type:        etype,
		Theme:       r.PostForm.Get("theme"),
		Description: r.PostForm.Get("description"),
		Genders:     genders,
		Age:         age,
		StartDate:   r.PostForm.Get("startDate"),
		EndDate:     r.PostForm.Get("endDate"),
	}, nil
}
