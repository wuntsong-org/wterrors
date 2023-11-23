package errors

const UnknownError = "UNKNOWN"

var BaseError = NewClass("base", "unknown error reason")

func Errorf(format string, a ...any) WTError {
	return BaseError.Errorf(format, a...)
}

func Warp(err error, format string, a ...any) WTError {
	return BaseError.Warp(err, format, a...)
}

func WarpQuick(err error) WTError {
	return BaseError.WarpQuick(err)
}
