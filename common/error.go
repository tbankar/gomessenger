package common

import "runtime/debug"

type ErrorMessage struct {
	LowLevel   error
	Message    string
	StackTrace string
	ExtraInfo  map[string]interface{}
}

func PopulateError(err error, msg string, st string, extra ...interface{}) ErrorMessage {
	return ErrorMessage{
		LowLevel:   err,
		Message:    msg,
		StackTrace: string(debug.Stack()),
		ExtraInfo:  make(map[string]interface{}),
	}
}
