package utils

import (
	"fmt"
	"net"
	"os"
	"strings"
	"time"
)

const host string = "localhost:8080"

func Server() {
	ln, err := net.Listen("tcp", host)
	handleError(err, "Failed to bind to port 8080!")

	for {
		// Awaiting for request from a client
		conn, err := ln.Accept()
		handleError(err, "Error accepting connection!")
		// Read request content from client
		b := make([]byte, 1024*4)
		conn.Read(b)

		lines := strings.Fields(string(b))
		_, path := lines[0], lines[1]

		// Determine response type based on request
		switch path {
		case "/":
			conn.Write(getHtml("./pages/html/index.html"))

		case "/api":
			resType := HelloWorld{Name: "Hello World!", Age: 69}
			res := JsonObject[HelloWorld]{Data: resType}
			conn.Write(res.GetJsonResponse("200 OK"))

		default:
			conn.Write(getError("404 Not Found", "./pages/html/errors/404.html"))
		}

		fmt.Println("Response was successful!")
	}
}

func getHtml(pageDir string) []byte {
	// Retrieve the page HTML from file
	page, err := os.ReadFile(pageDir)
	handleError(err, "Cannot extract content from HTML!")
	// Assemble the response headers
	head := Response{
		server:         host,
		date:           time.Now().UTC().Format(time.RFC1123),
		content_length: len(page),
		content_type:   "text/html",
		content:        page,
	}
	outStr := "HTTP/1.1 200 OK\r\n" +
		"Date: " + head.date + "\r\n" +
		"Content-Length: " + fmt.Sprint(head.content_length) + "\r\n" +
		"Content-Type: " + head.content_type + "\r\n\r\n"
	// Attach body to response headers and return as bytestring
	response := append([]byte(outStr), head.content...)
	return response
}

func getError(status_code string, pageDir string) []byte {
	// Retrieve the page HTML from file
	page, err := os.ReadFile(pageDir)
	handleError(err, "Cannot get error page!")
	// Assemble the response headers
	head := Response{
		server:         host,
		date:           time.Now().UTC().Format(time.RFC1123),
		content_length: len(page),
		content_type:   "text/html",
		content:        page,
	}
	outStr := "HTTP/1.1 " + status_code + "\r\n" +
		"Date: " + head.date + "\r\n" +
		"Content-Length: " + fmt.Sprint(head.content_length) + "\r\n" +
		"Content-Type: " + head.content_type + "\r\n\r\n"
	// Attach body to response headers and return as bytestring
	response := append([]byte(outStr), head.content...)
	return response
}

func (o *JsonObject[T]) GetJsonResponse(status_code string) []byte {
	body := o.GetJsonByteString()

	head := Response{
		server:         host,
		date:           time.Now().UTC().Format(time.RFC1123),
		content_length: len(body),
		content_type:   "application/json",
		content:        body,
	}
	outStr := "HTTP/1.1 " + status_code + "\r\n" +
		"Date: " + head.date + "\r\n" +
		"Content-Length: " + fmt.Sprint(head.content_length) + "\r\n" +
		"Content-Type: " + head.content_type + "\r\n\r\n"
	// Attach body to response headers and return as bytestring
	response := append([]byte(outStr), head.content...)
	return response
}
