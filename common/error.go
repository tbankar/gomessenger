package common

import "runtime/debug"

type ErrorMessage struct {
	LowLevel   error
	Message    string
	StackTrace string
	ExtraInfo  map[string]interface{}
}

//TODO: We have to log this error somewhere in a file
func PopulateError(err error, msg string, margs ...interface{}) ErrorMessage {
	return ErrorMessage{
		LowLevel:   err,
		Message:    msg,
		StackTrace: string(debug.Stack()),
		ExtraInfo:  make(map[string]interface{}),
	}
}
