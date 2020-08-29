package main

import (
	"log"
	"net/http"
)

func web(addr, path string) {

	// if our path is null we will return our in memory static
	// website!
	if path == "" {
		http.Handle("/", home)
	} else {
		http.Handle("/", http.FileServer(http.Dir(path)))
	}

	http.Handle("/ws", wserv)
	http.Handle("/api/health", health)
	http.Handle("/api/quote", theQuote)

	err := http.ListenAndServe(addr, nil)
	log.Fatal(err)
}
