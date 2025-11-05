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
	return fmt.Sprintf("%d", be.code)
}

func (be BusinessError) String() string {
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

const (
	ErrInternalCode             = 1000
	ErrUserNotFoundCode         = 1001
	ErrInvalidEventCode         = 1002
	ErrInvalidStartDateCode     = 1003
	ErrDocsNotFoundCode         = 1004
	ErrInvalidToGetEventsCode   = 1005
	ErrInvalidToCreateEventCode = 1006
	ErrInvalidToCreateIndexCode = 1007
	ErrInvalidToPublishMsgCode  = 1008
)

var (
	ErrInternal             = newBusinessError(ErrInternalCode, "internal business error")
	ErrUserNotFound         = newBusinessError(ErrUserNotFoundCode, "user not found")
	ErrInvalidEvent         = newBusinessError(ErrInvalidEventCode, "invalid event")
	ErrInvalidStartDate     = newBusinessError(ErrInvalidStartDateCode, "invalid start date")
	ErrDocsNotFound         = newBusinessError(ErrDocsNotFoundCode, "docs not found")
	ErrInvalidToGetEvents   = newBusinessError(ErrInvalidToGetEventsCode, "invalid to get events")
	ErrInvalidToCreateEvent = newBusinessError(ErrInvalidToCreateEventCode, "invalid to create event")
	ErrInvalidToCreateIndex = newBusinessError(ErrInvalidToCreateIndexCode, "invalid to create index")
	ErrInvalidToPublishMsg  = newBusinessError(ErrInvalidToPublishMsgCode, "invalid to publish msg")
)
