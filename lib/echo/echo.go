package main

import (
	"fmt"
	"net/http"
)

type Echo struct {
	Path string
}

var (
	Path    = "/api/echo"
	Service = &Echo{Path: Path}
)

func (e *Echo) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Foo")
}
