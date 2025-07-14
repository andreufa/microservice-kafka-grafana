package middleware

import "net/http"

type WripperWriter struct {
	http.ResponseWriter
	StatusCode int
}

func (w *WripperWriter) WriteHeader(statusCode int) {
	w.ResponseWriter.WriteHeader(statusCode)
	w.StatusCode = statusCode
}
