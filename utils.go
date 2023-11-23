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
		if isTarget {
			if wtTarget == wtErr {
				return true
			}

			if wtTarget.Code() == wtErr.Code() {
				return true
			}

			return errors.Is(err, wtTarget)
		}
	} else if errors.As(err, &normalErr) {
		return errors.Is(err, normalErr)
	}

	return false
}

func As(err error, target any) bool {
	return errors.As(err, target)
}
