package server

import (
	"net/http"
)

type Server struct {
	Port string
}

func (svr *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	svr.PrintRequestInfo(w, r)
	svr.PrintServerInfo(w, r)
}

// New returns a new server
func New(port string) http.Handler {
	return &Server{port}
}
