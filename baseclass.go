package errors

const UnknownError = "UNKNOWN"

var BaseError = newBaseClass("base", "unknown error reason")

func New(msgList ...string) WTError {
	return BaseError.New(msgList...)
}

func Errorf(format string, a ...any) WTError {
	return BaseError.Errorf(format, a...)
}

func Warp(err error, format string, a ...any) WTError {
	return BaseError.Warp(err, format, a...)
}

func WarpQuick(err error) WTError {
	return BaseError.WarpQuick(err)
}
