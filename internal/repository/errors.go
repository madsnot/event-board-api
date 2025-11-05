package repository

import (
	"fmt"
	"github.com/madsnot/event-board-api/internal/domain/models"
)

type IntegrationError struct {
	code int
	msg  string
}

func newIntegrationError(code int, msg string) IntegrationError {
	return IntegrationError{
		code: code,
		msg:  msg,
	}
}

func (ie IntegrationError) Error() string {
	return fmt.Sprintf("%d", ie.code)
}

func (ie IntegrationError) String() string {
	return fmt.Sprintf("%d: %s", ie.code, ie.msg)
}

func (ie IntegrationError) Code() int {
	return ie.code
}

func (ie IntegrationError) Msg() string {
	return ie.msg
}

func (ie IntegrationError) Wrap(err error) error {
	msg := ie.msg + ": " + err.Error()

	return models.NewWrapError([]int{ie.code}, []string{msg})
}

var (
	ErrInternal = newIntegrationError(2000, "internal integration error")
)
