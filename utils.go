package errors

import (
	"errors"
	"strings"
)

func NewClass(code string, msgList ...string) WTErrorClass {
	var msg string
	if len(msgList) == 0 {
		msg = strings.Replace(code, "_", " ", -1)
	} else {
		msg = msgList[0]
	}

	return &wtErrorClass{
		msg:  msg,
		code: code,
		base: false,
	}
}

func newBaseClass(code string, msgList ...string) WTErrorClass {
	var msg string
	if len(msgList) == 0 {
		msg = strings.Replace(code, "_", " ", -1)
	} else {
		msg = msgList[0]
	}

	return &wtErrorClass{
		msg:  msg,
		code: code,
		base: true,
	}
}

func Is(err error, target any) bool {
	var wtErr WTError
	var normalErr error

	if errors.As(err, &wtErr) {
		errClass, isClass := target.(WTErrorClass)
		if isClass {
			return errClass.Code() == wtErr.Code()
		}

		wtTarget, isTarget := target.(WTError)
		if !isTarget {
			return false
		}

		return errors.Is(err, wtTarget)
	} else if errors.As(err, &normalErr) {
		targetErr, targetIsErr := target.(error)
		if !targetIsErr {
			return false
		}

		return errors.Is(err, targetErr)
	}

	return false
}

func As(err error, target any) bool {
	return errors.As(err, target)
}
