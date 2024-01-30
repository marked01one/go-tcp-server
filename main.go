package main

import (
	"flag"
	"fmt"
	"tcp_server/server"
)

func main() {
	is_client := false
	flag.BoolVar(&is_client, "client", false, "determines whether the running TCP connection is a server or client.")

	flag.Parse()

	if !is_client {
		fmt.Print("::::: RUNNING SERVER :::::\n\n")
		server.Server()
	} else {
		fmt.Print("::::: RUNNING CLIENT :::::\n\n")
		server.Client()
	}
}
