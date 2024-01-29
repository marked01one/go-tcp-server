package utils

import "encoding/json"

type Response struct {
	server         string
	date           string
	content_length int
	content_type   string
	content        []byte
}

type JsonObject[T any] struct {
	Status string `json:"status"`
	Data   T      `json:"data"`
}

type HelloWorld struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func (o *JsonObject[T]) GetJsonByteString() []byte {
	j, err := json.MarshalIndent(o, "", "  ")
	handleError(err, "Cannot convert struct into JSON")
	return j
}
