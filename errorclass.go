package errors

import (
	"errors"
	"fmt"
)

type WTErrorClass interface {
	WTErrorClass()

	Code() string
	Msg() string

	New(...string) WTError
	Errorf(string, ...any) WTError
	Warp(error, string, ...any) WTError
	WarpQuick(error) WTError
}

type wtErrorClass struct {
	code string
	msg  string
	base bool
}

func (c *wtErrorClass) WTErrorClass() {}

func (c *wtErrorClass) Code() string {
	return c.code
}

func (c *wtErrorClass) Msg() string {
	return c.msg
}

func (c *wtErrorClass) New(msgList ...string) WTError {
	msg := c.msg
	if len(msgList) != 0 {
		msg = msgList[0]
	}

	code := c.code
	cause := error(nil)
	stack := getStack()

	return &wtError{
		class: c,
		cause: cause,
		msg:   msg,
		code:  code,
		stack: stack,
	}
}

func (c *wtErrorClass) Errorf(format string, a ...any) WTError {
	msg := fmt.Sprintf(format, a...)
	code := c.code
	cause := error(nil)
	stack := getStack()

	return &wtError{
		class: c,
		cause: cause,
		msg:   msg,
		code:  code,
		stack: stack,
	}
}

func (c *wtErrorClass) Warp(err error, format string, a ...any) WTError {
	if err == nil {
		return nil
	}

	msg := fmt.Sprintf(format, a...)
	code := c.code
	cause := err
	stack := ""

	var wtErr WTError
	if errors.As(err, &wtErr) {
		stack = wtErr.Stack()
	} else {
		stack = getStack()
	}

	return &wtError{
		class: c,
		cause: cause,
		msg:   msg,
		code:  code,
		stack: stack,
	}
}

func (c *wtErrorClass) WarpQuick(err error) WTError {
	if err == nil {
		return nil
	}

	var wtErr WTError
	if errors.As(err, &wtErr) && (c.base || wtErr.Code() == c.code) {
		return wtErr
	}

	msg := c.msg
	code := c.code
	cause := err
	stack := getStack()

	return &wtError{
		class: c,
		cause: cause,
		msg:   msg,
		code:  code,
		stack: stack,
	}
}
