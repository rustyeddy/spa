package main

import (
	"fmt"
	"log"
	"net/http"
)

var (
	health Health
	ws     wsServer
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
	err := http.ListenAndServe(addr, nil)
	log.Fatal(err)
}

// handleHome returns the html of our home page to the person making
// the request.
func handleHome(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	index := defaultIndex()
	fmt.Fprint(w, index)
}

// defaultIndex returns a very simple but complete static website
// as the index file for this program.
func defaultIndex() string {
	str := `
<!doctype html>
<html>
  <head>
    <title>Static</title>
  </head>
  <body>
    <h1>Hello, World!</h1>
    <p>This is a static page from memory.</p>
  </body>
</html>`
	return str
}
