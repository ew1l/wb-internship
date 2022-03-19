package server

import (
	"log"
	"net/http"
)

// StatusRecorder for getting the status code
type StatusRecorder struct {
	http.ResponseWriter
	Status int
}

// WriteHeader to send status code
func (sr *StatusRecorder) WriteHeader(status int) {
	sr.ResponseWriter.WriteHeader(status)
	sr.Status = status
}

// WithLogging for middleware
func WithLogging(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		recorder := &StatusRecorder{
			ResponseWriter: w,
			Status:         http.StatusOK,
		}

		h.ServeHTTP(recorder, r)
		log.Printf("%s %s %d", r.Method, r.RequestURI, recorder.Status)
	})
}
