package main

import (
	"flag"
	"log"
)

type Configuration struct {
	Addr  string // Address we will serve from
	Debug bool
	Pub   string // path we will publish
}

var (
	config Configuration
)

func init() {
	flag.BoolVar(&config.Debug, "debug", false, "turn on debugging")
	flag.StringVar(&config.Addr, "addr", ":8080", "address string to serve up")
	flag.StringVar(&config.Pub, "pub", "www", "path of website to serve up")
}

func main() {
	flag.Parse()

	// Let the world know we are alive
	log.Println("Web starting at", config.Addr)

	// Go web! Found in web.go
	web(config.Addr, config.Pub)
}
