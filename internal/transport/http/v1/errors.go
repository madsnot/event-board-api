package v1

import (
	"fmt"
	"github.com/madsnot/event-board-api/internal/domain/models"
	"net/http"
	"strconv"
	"strings"
)

type TransportError struct {
	code int
	msg  string
}

func newTransportError(code int, msg string) TransportError {
	return TransportError{
		code: code,
		msg:  msg,
	}
}

func (te TransportError) Error() string {
	return fmt.Sprintf("%d: %s", te.code, te.msg)
}

func (te TransportError) Code() int {
	return te.code
}

func (te TransportError) Msg() string {
	return te.msg
}

func (te TransportError) Wrap(err error) error {
	codesStr := strings.Fields(err.Error())

	codes := make([]int, 1, len(codesStr)+1)

	codes[0] = te.code

	for _, str := range codesStr {
		code, _ := strconv.Atoi(str)
		codes = append(codes, code)
	}

	return models.NewWrapError(codes)
}

var (
	ErrBadRequest = newTransportError(http.StatusBadRequest, "bad request")
	ErrInternal   = newTransportError(http.StatusInternalServerError, "internal error")
)
