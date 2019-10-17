package main

import (
	"net/http"
	"time"
)

// NewServer creates new server
func NewServer() *http.Server {

	srv := http.Server{
		Addr:         ":8080",
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		Handler:      rateLimit(http.HandlerFunc(WebconvHandler)),
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
