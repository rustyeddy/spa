package main

import (
	"flag"
	"log"
)

type Configuration struct {
	Addr string // Address we will serve from
	Pub  string // path we will publish
	Lib  string
}

var (
	config Configuration
)

func init() {
	flag.StringVar(&config.Addr, "addr", ":1233", "address string to serve up")
	flag.StringVar(&config.Pub, "pub", "", "path of website to serve up")
	flag.StringVar(&config.Lib, "lib", "../lib", "The library")
}

func main() {
	flag.Parse()

	// Read plugins
	plugins(config.Lib)

	// Let the world know we are alive
	log.Println("Web starting at", config.Addr)

	// Go web! Found in web.go
	web(config.Addr, config.Pub)
}
