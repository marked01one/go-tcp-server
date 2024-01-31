package server

import (
	"fmt"
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
		fmt.Println(string(b))
		lines := strings.Fields(string(b))
		req := Request{}
		req.Parse(lines)

		// Determine response type based on request
		switch req.path {

		case "/":
			conn.Write(getPage("200 OK", "./pages/html/index.html"))
			logRes(req.method, req.path, "200 OK")

		case "/api":
			if req.method == "GET" {
				helloWorld := HelloWorld{Name: "Hello World!", Age: 69}
				res := JsonObject[HelloWorld]{Data: helloWorld}
				conn.Write(res.GetJsonResponse("200 OK"))
				logRes(req.method, req.path, "200 OK")

			} else if req.method == "POST" {
				// Handle case where response body is not in JSON
				if req.content_type != "application/json" {
					conn.Write(getBasicError("400 Bad Request"))
					break
				}
				// Generate JSON response body
				helloWorld := HelloWorld{Name: "Hello World!", Age: 69}
				res := JsonObject[HelloWorld]{Data: helloWorld}
				conn.Write(res.GetJsonResponse("200 OK"))
				logRes(req.method, req.path, "200 OK")

			}

		default:
			conn.Write(getPage("404 Not Found", "./pages/html/errors/404.html"))
			logRes(req.method, req.path, "404 Not Found")

		}

		conn.Close()
	}
}
