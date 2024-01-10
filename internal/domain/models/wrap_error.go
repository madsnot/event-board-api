package models

import "strconv"

type WrapError struct {
	codes []int
	msgs  []string
}

func NewWrapError(codes []int, msgs []string) WrapError {
	return WrapError{
		codes: codes,
		msgs:  msgs,
	}
}

func (we WrapError) Codes() []int {
	return we.codes
}

func (we WrapError) Msgs() []string {
	return we.msgs
}

func (we WrapError) String() string {
	var errStr string

	for ind, code := range we.codes {
		errStr += strconv.Itoa(code) + we.msgs[ind]

		if ind < len(we.codes) {
			errStr += "; "
		}
	}

	return errStr
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
