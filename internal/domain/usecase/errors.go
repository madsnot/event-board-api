package usecase

import (
	"fmt"
	"github.com/madsnot/event-board-api/internal/domain/models"
	"strconv"
	"strings"
)

type BusinessError struct {
	code int
	msg  string
}

func newBusinessError(code int, msg string) BusinessError {
	return BusinessError{
		code: code,
		msg:  msg,
	}
}

func (be BusinessError) Error() string {
	return fmt.Sprintf("%d: %s", be.code, be.msg)
}

func (be BusinessError) Code() int {
	return be.code
}

func (be BusinessError) Msg() string {
	return be.msg
}

func (be BusinessError) Wrap(err error) error {
	codesStr := strings.Fields(err.Error())

	codes := make([]int, 1, len(codesStr)+1)

	codes[0] = be.code

	for _, str := range codesStr {
		code, _ := strconv.Atoi(str)
		codes = append(codes, code)
	}

	return models.NewWrapError(codes)
}

var (
	ErrUserNotFound     = newBusinessError(1000, "user not found")
	ErrInvalidEvent     = newBusinessError(1001, "invalid event")
	ErrInvalidStartDate = newBusinessError(1002, "invalid start date")
)
