package utils

import (
	"fmt"
	"net"
	"os"
	"strings"
)

type request struct {
	method string
	path   string
}

func Server() {
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("Failed to bind to port 8080: ", err.Error())
		os.Exit(1)
	}

	conn, err := ln.Accept()
	if err != nil {
		fmt.Println("Error accepting connection: ", err.Error())
		os.Exit(1)
	} else {
		fmt.Println("Connection Accepted!")
	}

	b := make([]byte, 1024)

	conn.Read(b)
	lines := strings.Fields(string(b))

	req := request{
		method: lines[0],
		path:   lines[1],
	}
	fmt.Print(string(b))

	if req.path == "/" {
		conn.Write([]byte("HTTP/1.1 200 OK\r\n\r\n"))
	} else {
		conn.Write([]byte("HTTP/1.1 404 Not Found\r\n\r\n"))
	}

	conn.Close()
}
