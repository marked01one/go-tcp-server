package server

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func Client() {

	for {
		conn, _ := net.Dial("tcp", ":8080")

		hi, err := input("> ")
		if err != nil {
			fmt.Println(err.Error())
		}

		conn.Write([]byte(hi))
	}
}

type Input struct {
}

func input(prompt string) (string, error) {
	fmt.Print(prompt)
	reader := bufio.NewReader(os.Stdin)
	msg, _, err := reader.ReadLine()
	return string(msg), err
}
