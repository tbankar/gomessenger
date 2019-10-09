package common

import "net/http"

func ResponseToClient(statusCode int, msg string, w http.ResponseWriter) {
	w.WriteHeader(statusCode)
	w.Write([]byte(msg))
	return
}
