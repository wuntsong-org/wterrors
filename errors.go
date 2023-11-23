package errors

import (
	"errors"
	"fmt"
)

const UnknownError = "UNKNOWN"

type WTError interface {
	/* 标记类操作 */
	error
	WTError() // 标记

	/* 访问类操作 */
	Code() string
	Message() string
	Stack() string
	Cause() error

	/* 设置类操作 */
	SetCode(code string) WTError
	SetCodeForce(code string) WTError
	SetCause(cause error) WTError
	SetCauseForce(cause error) WTError
	Warp(format string, a ...any) WTError
}

type wtError struct {
	cause error
	msg   string
	code  string
	stack string
}

func (*wtError) WTError() {}

func (w *wtError) Code() string {
	return w.code
}

func (w *wtError) Message() string {
	if w.cause == nil {
		return fmt.Sprintf("[%s] %s", w.code, w.msg)
	}

	var cause *wtError
	ok := errors.As(w.cause, &cause)
	if ok {
		return fmt.Sprintf("[%s] %s: %s", w.code, w.msg, cause.Message())
	}

	return fmt.Sprintf("[%s] %s: %s", w.code, w.msg, w.cause.Error())
}

func (w *wtError) Stack() string {
	return w.stack
}

func (w *wtError) Error() string {
	return w.Message()
}

func (w *wtError) Cause() error {
	return w.cause
}

func (w *wtError) SetCode(code string) WTError {
	if w.code == UnknownError {
		w.code = code
	}
	return w
}

func (w *wtError) SetCodeForce(code string) WTError {
	w.code = code
	return w
}

func (w *wtError) SetCause(cause error) WTError {
	if w.cause == nil {
		w.cause = cause
		_ = w.SetCode(getErrorName(cause))
	}
	return w
}

func (w *wtError) SetCauseForce(cause error) WTError {
	w.cause = cause
	_ = w.SetCode(getErrorName(cause))
	return w
}

func (w *wtError) Warp(format string, a ...any) WTError {
	return Warp(w, format, a...)
}
