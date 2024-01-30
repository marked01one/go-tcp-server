package server

import (
	"net"
	"strings"
)

func Server() {
	ln, err := net.Listen("tcp", HOST)
	panicErr(err, "Failed to bind to port 8080!")

	for {
		// Awaiting for request from a client
		conn, err := ln.Accept()
		panicErr(err, "Error accepting connection!")
		// Read request content from client
		b := make([]byte, 1024*4)
		conn.Read(b)

		lines := strings.Fields(string(b))
		method, path := lines[0], lines[1]

		// Determine response type based on request
		switch path {

		case "/":
			conn.Write(getPage("200 OK", "./pages/html/index.html"))
			logRes(method, path, "200 OK")

		case "/api":
			if method == "GET" {
				helloWorld := HelloWorld{Name: "Hello World!", Age: 69}
				res := JsonObject[HelloWorld]{Data: helloWorld}
				conn.Write(res.GetJsonResponse("200 OK"))
				logRes(method, path, "200 OK")

			} else if method == "POST" {
				continue
			}

		default:
			conn.Write(getPage("404 Not Found", "./pages/html/errors/404.html"))
			logRes(method, path, "404 Not Found")

		}

	}
}
