package main

import (
	"fmt"
	"net/http"
)

var (
	home Home
)

type Home struct {
}

// handleHome returns the html of our home page to the person making
// the request.
func (h Home) ServeHTTP(w http.ResponseWriter, r *http.Request) {
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
    <h1>Hello, Static Website!</h1>
    <p>This is a static page straight from memory.</p>
  </body>
</html>`
	return str
}
