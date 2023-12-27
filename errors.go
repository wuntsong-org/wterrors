package errors

import (
	"fmt"
)

type WTError interface {
	/* 标记类操作 */
	error
	WTError_TAG() // 标记

	/* 访问类操作 */
	Class() WTErrorClass
	Code() string
	Msg() string
	Message() string
	MessageWithStack() string
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
	class WTErrorClass
}

func (*wtError) WTError_TAG() {}

func (w *wtError) Class() WTErrorClass {
	return w.class
}

func (w *wtError) Code() string {
	return w.code
}

func (w *wtError) Msg() string {
	return w.msg
}

func (w *wtError) Message() string {
	if w.cause == nil {
		return fmt.Sprintf("[%s] %s", w.code, w.msg)
	}

	var cause *wtError
	ok := As(w.cause, &cause)
	if ok {
		return fmt.Sprintf("[%s] %s: %s", w.code, w.msg, cause.Message())
	}

	return fmt.Sprintf("[%s] %s: %s", w.code, w.msg, w.cause.Error())
}

func (w *wtError) MessageWithStack() string {
	if w.cause == nil {
		return fmt.Sprintf("[%s]%s\n%s", w.code, w.msg, w.stack)
	}

	var cause *wtError
	ok := As(w.cause, &cause)
	if ok {
		return fmt.Sprintf("[%s]%s: %s\n%s", w.code, w.msg, cause.Message(), w.stack)
	}

	return fmt.Sprintf("[%s]%s: %s\n%s", w.code, w.msg, w.cause.Error(), w.stack)
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
		return w.SetCodeForce(code)
	}
	return w
}

func (w *wtError) SetCodeForce(code string) WTError {
	w.code = code
	return w
}

func (w *wtError) SetCause(cause error) WTError {
	if w.cause == nil {
		return w.SetCauseForce(cause)
	}
	return w
}

func (w *wtError) SetCauseForce(cause error) WTError {
	w.cause = cause
	_ = w.SetCode(getErrorName(cause))

	var wtErr WTError
	if As(cause, &wtErr) {
		w.stack = wtErr.Stack()
	}

	return w
}

func (w *wtError) Warp(format string, a ...any) WTError {
	return Warp(w, format, a...)
}
