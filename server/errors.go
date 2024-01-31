package server

import "time"

func getBasicError(status_code string) []byte {
	outStr := "HTTP/1.1 " + status_code + "\r\n" +
		"Date: " + time.Now().UTC().Format(time.RFC1123) + "\r\n\r\n"
	// Attach body to response headers and return as bytestring
	response := []byte(outStr)
	return response
}
