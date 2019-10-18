package main

import (
	"net/http"
	"time"

	"github.com/didip/tollbooth"
	"github.com/didip/tollbooth/limiter"
)

// NewServer creates new server and limiter
func NewServer() *http.Server {

	lmt := tollbooth.NewLimiter(500, &limiter.ExpirableOptions{DefaultExpirationTTL: 5 * time.Second})
	lmt.SetIPLookups([]string{"X-Forwarded-For", "RemoteAddr", "X-Real-IP"})

	h := http.HandlerFunc(WebconvHandler)

	srv := http.Server{
		Addr:         ":8080",
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		Handler:      tollbooth.LimitFuncHandler(lmt, h), // handle with third-party limiter
	}

	srv.SetKeepAlivesEnabled(false)

	return &srv
}

// WebconvHandler a http.handler function
func WebconvHandler(w http.ResponseWriter, r *http.Request) {

	conv := NewConv()
	conv.CopyInput(r)
	err := conv.SwapFormat()
	conv.MakeReply(w, err)

}
