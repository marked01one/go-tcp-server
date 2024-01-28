package main

import (
	"flag"
	"fmt"
	"tcp_server/utils"
)

func main() {
	is_server := false
	flag.BoolVar(&is_server, "server", false, "determines whether the running TCP connection is a server or client.")

	flag.Parse()

	fmt.Println(utils.Greeting())
}
