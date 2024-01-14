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
	msgs := make([]string, 1, len(codesStr)+1)

	codes[0] = be.code
	msgs[0] = be.msg

	for ind, str := range codesStr {
		if ind%2 != 0 {
			code, _ := strconv.Atoi(str)

			codes = append(codes, code)
		} else {
			msg := str

			msgs = append(msgs, msg)
		}

	}

	return models.NewWrapError(codes, msgs)
}

var (
	ErrInternal             = newBusinessError(1000, "internal business error")
	ErrUserNotFound         = newBusinessError(1001, "user not found")
	ErrInvalidEvent         = newBusinessError(1002, "invalid event")
	ErrInvalidStartDate     = newBusinessError(1003, "invalid start date")
	ErrDocsNotFound         = newBusinessError(1004, "docs not found")
	ErrInvalidToGetEvents   = newBusinessError(1005, "invalid to get events")
	ErrInvalidToCreateEvent = newBusinessError(1006, "invalid to create event")
	ErrInvalidToCreateIndex = newBusinessError(1007, "invalid to create index")
	ErrInvalidToPublishMsg  = newBusinessError(1008, "invalid to publish msg")
)
