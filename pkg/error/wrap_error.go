package error

import (
	"strconv"
	"strings"
)

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

	return NewWrapError(codes, msgs)
}
