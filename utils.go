package errors

import (
	"errors"
	"fmt"
)

func Errorf(format string, a ...any) WTError {
	msg := fmt.Sprintf(format, a...)
	code := UnknownError
	cause := error(nil)
	stack := getStack()

	return &wtError{
		cause: cause,
		msg:   msg,
		code:  code,
		stack: stack,
	}
}

func Warp(err error, format string, a ...any) WTError {
	msg := fmt.Sprintf(format, a...)
	code := getErrorName(err)
	cause := err
	stack := getStack()

	return &wtError{
		cause: cause,
		msg:   msg,
		code:  code,
		stack: stack,
	}
}

func WarpQuick(err error) WTError {
	msg := "error"
	code := getErrorName(err)
	cause := err
	stack := getStack()

	return &wtError{
		cause: cause,
		msg:   msg,
		code:  code,
		stack: stack,
	}
}

func Is(err error, target error) bool {
	var wtErr, wtTarget WTError

	if errors.As(err, &wtErr) && errors.As(target, &wtTarget) {
		if wtErr.Code() == wtTarget.Code() {
			return true
		} else if wtErr == wtTarget {
			return true
		}
	}

	return errors.Is(err, target)
}

func As(err error, target any) bool {
	return errors.As(err, target)
}
