package server

import (
	"fmt"
	"os"
	"time"
)

func getPage(status_code string, pageDir string) []byte {
	// Retrieve the page HTML from file
	page, err := os.ReadFile(pageDir)
	panicErr(err, "Cannot get HTML page content!")
	// Assemble the response headers
	head := Response{
		server:         HOST,
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
