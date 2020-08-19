package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"plugin"
)

type Server struct {
	Addr string
	Path string
}

func NewService(addr, path string) (s *Service) {
	s = &Service{addr, path}

	return s
}

func (s *Service) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println("This is service.ServeHTTP")
}

// plugins
func plugins(opath string) {

	//list := make([]string, 0, 10)
	//dirs := make([]string, 0, 10)

	err := filepath.Walk(opath, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		if filepath.Ext(path) == ".so" {
			openPlugin(path)
		}
		return nil
	})
	if err != nil {
		log.Fatal("error: ", opath)
	}
}

func openPlugin(path string) {
	plug, err := plugin.Open(path)
	if err != nil {
		log.Fatal("error: ", err)
	}

	s, err := plug.Lookup("Service")
	if err != nil {
		panic(err)
	}

	p, err := plug.Lookup("Path")
	if err != nil {
		panic(err)
	}

	path = *p.(*string)
	sf := s.(http.Handler)

	fmt.Printf("p: %s\n", path)
	fmt.Printf("s: %+v\n", sf)

	//http.Handle(path, sf)
}
