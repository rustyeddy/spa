package main

import (
	"flag"
	"fmt"
)

type Configuration struct {
	Addr string // Address we will serve from
	Pub  string // path we will publish
}

var (
	config Configuration
)

func init() {
	flag.StringVar(&config.Addr, "addr", ":1233", "address string to serve up")
	flag.StringVar(&config.Pub, "pub", "", "path of website to serve up")
}

func main() {
	flag.Parse()

	// Let the world know we are alive
	fmt.Println("Web starting at", config.Addr)

	// Go web!
	web(config.Addr, config.Pub)
}
