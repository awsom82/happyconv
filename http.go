package main

import (
	"log"
	"net/http"
)

type Server struct {
	TotalRequests int // number of all requests since start
}

func NewServer() *Server {
	return &Server{}
}

func (srv *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	srv.TotalRequests++

	conv := NewConv()
	conv.CopyInput(r)
	conv.SwapFormat()
	conv.MakeReply(w)

	log.Printf("№%d request served\n", srv.TotalRequests)

}
