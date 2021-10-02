package net

import (
	"log"
	"net/http"
)

type Server struct {
	Handlers	[]string
}

func init() {
}

// Register end points and an object that has a ServeHTTP method
func (s *Server) Register(path string, handler http.Handler) {
	s.Handlers = append(s.Handlers, path)
	http.Handle(path, handler)
}

// Run the web server for HTML in path under the given address
// string example "0.0.0.0:8080"
func (s *Server) Run(addr, path string) {
	err := http.ListenAndServe(addr, nil)
	log.Fatal(err)
}
