package models

import "strconv"

type WrapError struct {
	codes []int
}

func NewWrapError(codes []int) WrapError {
	return WrapError{
		codes: codes,
	}
}

func (we WrapError) Codes() []int {
	return we.codes
}

func (we WrapError) Error() string {
	var errStr string

	for ind, code := range we.codes {
		errStr += strconv.Itoa(code)

		if ind < len(we.codes) {
			errStr += " "
		}
	}

	return errStr
}
