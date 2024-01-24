package main

import (
	"bufio"
	"flag"
	"fmt"
	"net"
	"os"
)

func main() {
	is_server := false
	flag.BoolVar(&is_server, "server", false, "determines whether the running TCP connection is a server or client.")

	flag.Parse()

	if is_server {
		server()
	} else {
		client()
	}
}

func client() {
	fmt.Println(":::: RUNNING CLIENT ::::")
	scanner := bufio.NewReader(os.Stdin)

	for {
		// Receive string from user
		fmt.Print("> ")
		msg, _, err := scanner.ReadLine()
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		// Send message to server
		conn, _ := net.Dial("tcp", ":8080")
		conn.Write(msg)

		// Waits to receive response from server
		b := make([]byte, 1024)
		conn.Read(b)
		fmt.Println("[res] ", string(b))
	}

}

func server() {
	listener, err := net.Listen("tcp", ":8080")
	fmt.Println(":::: RUNNING SERVER ::::")

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	for {
		// Receive and print message from client
		connServer, _ := listener.Accept()
		b := make([]byte, 1024)
		connServer.Read(b)
		fmt.Println("[Client] ", string(b))

		// Prepare response message
		connServer.Write([]byte("re: Message sent successfully!"))
	}

}
