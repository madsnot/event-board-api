package v1

import (
	"github.com/madsnot/event-board-api/internal/domain/usecase"
	"net/http"
)

var (
	CODES = map[int]int{
		usecase.ErrUserNotFoundCode:         http.StatusNotFound,
		usecase.ErrInvalidEventCode:         http.StatusBadRequest,
		usecase.ErrInvalidStartDateCode:     http.StatusBadRequest,
		usecase.ErrDocsNotFoundCode:         http.StatusNotFound,
		usecase.ErrInvalidToGetEventsCode:   http.StatusInternalServerError,
		usecase.ErrInvalidToCreateEventCode: http.StatusInternalServerError,
		usecase.ErrInvalidToCreateIndexCode: http.StatusInternalServerError,
		usecase.ErrInvalidToPublishMsgCode:  http.StatusInternalServerError,
	}
)
