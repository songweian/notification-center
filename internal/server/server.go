package server

import (
	"net/http"
)

type Server struct {
	// ...
}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) HandleRequest(w http.ResponseWriter, r *http.Request) {
	// ...
}
