package utils

import (
	"fmt"
	"net"
	"os"
	"strings"
	"time"
)

type request struct {
	method string
	path   string
}

type Response struct {
	server         string
	date           string
	content_length int
	content_type   string
	content        []byte
}

const host string = "localhost:8080"

func Server() {
	ln, err := net.Listen("tcp", host)
	handleError(err, "Failed to bind to port 8080!")
	// Awaiting for request from a client
	conn, err := ln.Accept()
	handleError(err, "Error accepting connection!")
	fmt.Println("Connection Accepted!")
	// Read request content from client
	b := make([]byte, 1024*4)
	conn.Read(b)
	lines := strings.Fields(string(b))
	req := request{method: lines[0], path: lines[1]}
	// Determine response type based on request
	switch req.path {
	case "/":
		getPage(conn, "./pages/html/index.html")
	default:
		conn.Write([]byte("HTTP/1.1 404 Not Found\r\n\r\n"))
	}

	fmt.Println("Response was successful!")
	conn.Close()
}

func getPage(conn net.Conn, pageDir string) {
	page, err := os.ReadFile(pageDir)
	handleError(err, "Cannot extract content from HTML!")

	head := Response{
		server:         host,
		date:           time.Now().UTC().Format(time.RFC1123),
		content_length: len(page),
		content_type:   "text/html",
		content:        page,
	}

	response := head.generateResponse("HTTP/1.1 200 OK")

	conn.Write(response)
}

func (h *Response) generateResponse(response_code string) []byte {
	outStr := response_code + "\r\n" +
		"Date: " + h.date + "\r\n" +
		"Content-Length: " + fmt.Sprint(h.content_length) + "\r\n" +
		"Content-Type: " + h.content_type + "\r\n\r\n"

	return append([]byte(outStr), h.content...)
}
