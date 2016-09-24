package main

import "flag"

var (
	host = flag.String("host", "0.0.0.0", "default host")
	port = flag.Int("port", 8080, "default port")
)

func main() {
	flag.Parse()

	go startDefaultEp(*host, *port)
	startShell()
}
