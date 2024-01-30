package server

import (
	"encoding/json"
	"fmt"
	"time"
)

type JsonObject[T any] struct {
	Status string `json:"status"`
	Data   T      `json:"data"`
}

func (o *JsonObject[T]) getJsonByteString() []byte {
	j, err := json.MarshalIndent(o, "", "  ")
	panicErr(err, "Cannot convert struct into JSON")
	return j
}

func (o *JsonObject[T]) GetJsonResponse(status_code string) []byte {
	body := o.getJsonByteString()

	head := Response{
		server:         HOST,
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

type HelloWorld struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type MRA struct {
}
