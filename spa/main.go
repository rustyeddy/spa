package main

import (
	"fmt"
	"net/http"
)

func main() {

	// Let the world know we are alive
	fmt.Println("Hello, world!")

	http.HandleFunc("/", handleHome)
	//http.Handle("/api/health", handleHealth)
	//http.Handle("/", http.FileServer(http.Dir("/pub")))

	http.ListenAndServe(":1233", nil)
}

// handleHome returns the html of our home page to the person making
// the request.
func handleHome(w http.ResponseWriter, r *http.Request) {

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
    <p>This is an extremly simple static web page (with bad spelling)</p>
  </body>
</html>`
	return str
}
