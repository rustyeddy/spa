package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Health struct {
	Health string `json:"health"`
}

func web(addr, path string) {

	// if our path is null we will return our in memory static
	// website!
	if path == "" {
		http.HandleFunc("/", handleHome)
	} else {
		http.Handle("/", http.FileServer(http.Dir(path)))
	}

	http.HandleFunc("/api/health", handleHealth)
	http.ListenAndServe(addr, nil)
}

// handleHome returns the html of our home page to the person making
// the request.
func handleHome(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	index := defaultIndex()
	fmt.Fprint(w, index)
}

func handleHealth(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Health{"ok"})
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
