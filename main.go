package main

import (
	"flag"
	"fmt"
	"net"
)

func client() {
	conn, _ := net.Dial("tcp", ":8080")

	conn.Write([]byte("Hello World!"))
}

func server() {
	listener, _ := net.Listen("tcp", ":8080")

	b := make([]byte, 1024)

	for {
		conn, _ := listener.Accept()
		_, _ = conn.Read(b)
		fmt.Println("[Client]\t", string(b))
	}

}

func main() {
	is_server := false
	flag.BoolVar(&is_server, "server", false, "determines whether the running TCP connection is a server or client.")

	flag.Parse()

	if is_server {
		fmt.Println(":::: RUNNING SERVER ::::")
		server()
	} else {
		fmt.Println(":::: RUNNING CLIENT ::::")
		client()
	}
}
