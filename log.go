package webconv

import (
	"log"
	"net/http"
)

// Log colors
const (
	InfoColor    = "\033[1;34m%s\033[0m %s"
	WarningColor = "\033[1;33m%s\033[0m %s"
)

// WebconvLogMiddleware logs http requests
func WebconvLogMiddleware(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		color := InfoColor

		if r.Method != "POST" {
			color = WarningColor
		}

		log.Printf(color, r.Method, r.URL.Path)

		next.ServeHTTP(w, r)
	})
}
