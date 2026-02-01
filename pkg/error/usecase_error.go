package error

import (
	"fmt"
	"strconv"
	"strings"
)

type BusinessError struct {
	code int
	msg  string
}

func NewBusinessError(code int, msg string) BusinessError {
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

	return NewWrapError(codes, msgs)
}
