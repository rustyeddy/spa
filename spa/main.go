package main

import (
	"flag"
	"log"
	"os"

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
}

func main() {
	flag.Parse()

	if len(os.Args) > 1 {
		config.Pub = os.Args[1]
		log.Println("HTML directory has been set to ", config.Pub)
	}



	addRoutes()

	// Let the world know we are alive
	log.Println("Web starting at", config.Addr)
	// Go web! Found in web.go
	log.Println("start web pages at: ", config.Pub)
	srv.Run(config.Addr, config.Pub)
}


func addRoutes() {

	// Register our REST handler callbacks. For google appengine it needs to
	// be done here in the init function to ensure the callbacks have been
	// registered in other modules, if they exist.
	if config.Pub == "" {
		srv.Register("/", home)
	} else {
		srv.Register("/", http.FileServer(http.Dir(config.Pub)))
	}

	srv.Register("/ws", wserv)
	srv.Register("/api/health", health)
	srv.Register("/api/quote", quote)
}
