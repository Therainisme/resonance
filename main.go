package main

import (
	"flag"
	"resonance/client"
	"resonance/server"
)

func main() {
	flag.Parse()

	if *server.MusicName != "" {
		server.Run()
	} else {
		client.Run()
	}
}
