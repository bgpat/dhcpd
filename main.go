package main

import (
	"log"

	"github.com/bgpat/dhcpd/server"
)

func main() {
	s, err := server.New()
	if err == nil {
		err = s.Listen()
	}
	log.Fatal(err.Error())
}
