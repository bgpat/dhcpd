package main

import (
	"fmt"
	"os"

	"github.com/bgpat/dhcpd/server"
)

func main() {
	s, err := server.New(func(lease *server.Lease) server.Reply {
		fmt.Printf("lease: %+v\n", lease)
		return nil
	})
	if err == nil {
		fmt.Printf("%+v\n", s.Handler)
		err = s.Listen()
	}
	fmt.Fprintln(os.Stderr, err.Error())
}
