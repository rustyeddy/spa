package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Hello, world!")

	http.HandleFunc("/", handleHome)
	//http.Handle("/api/health", handleHealth)
	//http.Handle("/", http.FileServer(http.Dir("/pub")))

	http.ListenAndServe(":1233", nil)
}

func handleHome(w http.ResponseWriter, r *http.Request) {

	index := defaultIndex()
	fmt.Fprint(w, index)
}

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
