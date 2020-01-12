package middlewares

import (
	"net/http"
	"net/http/httptest"

	log "github.com/sirupsen/logrus"
)

// JSONMiddleware sets JSON as default content type.
func JSONMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-type", "application/json")
		next.ServeHTTP(w, r)
	})
}

// LoggingMiddleware logs each request and response
func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Infof("%s %s | %s %d", r.Method, r.URL.Path, r.Proto, r.ContentLength)
		rec := httptest.NewRecorder()
		next.ServeHTTP(rec, r)

		for k, v := range rec.Header() {
			w.Header()[k] = v
		}

		log.Infof("%s %s (%d) | %s %d", r.Method, r.URL.Path, rec.Result().StatusCode,
			r.Proto, rec.Result().ContentLength)

		w.Write(rec.Body.Bytes())
	})
}
