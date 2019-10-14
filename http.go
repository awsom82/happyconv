package main

import (
	"net/http"
	"time"
)

// Creates new server
func NewServer() *http.Server {
	srv := http.Server{
		Addr:         ":8080",
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		Handler:      http.HandlerFunc(WebconvHadler),
	}

	srv.SetKeepAlivesEnabled(false)

	return &srv
}

// http.handler
func WebconvHadler(w http.ResponseWriter, r *http.Request) {

	conv := NewConv()
	conv.CopyInput(r)
	err := conv.SwapFormat()
	conv.MakeReply(w, err)

}
