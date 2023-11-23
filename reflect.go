package errors

import "reflect"

func getErrorName(err error) string {
	t := reflect.TypeOf(err)
	if t.Kind() == reflect.Invalid {
		return UnknownError
	}

	for {
		if t.Kind() == reflect.Ptr || t.Kind() == reflect.Interface {
			t = t.Elem()
		} else {
			return t.Name()
		}
	}
}
