package middleware

import (
	"filter-service/metrics"
	"log"
	"net/http"
	"time"
)

func RawLog(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()
			wrapper := &WripperWriter{
				ResponseWriter: w,
				StatusCode:     http.StatusOK,
			}
			next.ServeHTTP(wrapper, r)
			metrics.ObserveRequest(r.Method, "Raw", wrapper.StatusCode)
			log.Println(wrapper.StatusCode, r.Method, r.URL.Path, time.Since(start))
		})
}
