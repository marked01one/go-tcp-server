package main

import (
	"flag"
	"fmt"
	"tcp_server/utils"
)

func main() {
	is_client := false
	flag.BoolVar(&is_client, "client", false, "determines whether the running TCP connection is a server or client.")

	flag.Parse()

	if !is_client {
		fmt.Print("::::: RUNNING SERVER :::::\n\n")
		utils.Server()
	} else {
		fmt.Print("::::: RUNNING CLIENT :::::\n\n")
		utils.Client()
	}
}
