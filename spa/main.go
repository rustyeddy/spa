package main

import (
	"flag"
	"log"

	"net/http"

	"github.com/rustyeddy/spa/net"
	"github.com/rustyeddy/spa/handlers"
)

type Configuration struct {
	Addr  string // Address we will serve from
	Debug bool
	Pub   string // path we will publish
}

var (
	config Configuration
	srv	net.Server

	home handlers.Home
	health handlers.Health
	quote handlers.Quote
	wserv net.WebSocket
)

func init() {
	flag.BoolVar(&config.Debug, "debug", false, "turn on debugging")
	flag.StringVar(&config.Addr, "addr", ":8080", "address string to serve up")
	flag.StringVar(&config.Pub, "pub", "www", "path of website to serve up")

	// Register our REST handler callbacks
	if config.Pub == "" {
		srv.Register("/", home)
	} else {
		srv.Register("/", http.FileServer(http.Dir(config.Pub)))
	}

	srv.Register("/ws", wserv)
	srv.Register("/api/health", health)
	srv.Register("/api/quote", quote)
}

func main() {
	flag.Parse()

	// Let the world know we are alive
	log.Println("Web starting at", config.Addr)

	// Go web! Found in web.go
	srv.Run(config.Addr, config.Pub)
}
