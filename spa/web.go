package main

import (
	"log"
	"net/http"
)

var (
	health Health
	quote  Quote
)

type Server struct {
}

func web(addr, path string) {

	// if our path is null we will return our in memory static
	// website!
	if path == "" {
		http.HandleFunc("/", handleHome)
	} else {
		http.Handle("/", http.FileServer(http.Dir(path)))
	}

	http.Handle("/ws", ws)
	http.Handle("/api/health", health)
	http.Handle("/api/quote", quote)

	err := http.ListenAndServe(addr, nil)
	log.Fatal(err)
}
