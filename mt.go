package main

import (
	"fmt"
	"net"
	"os"
)

const PORT = "2525"

func main() {
	server := NewTCPServer(PORT)

	defer server.Close()

	fmt.Printf("Listening on port %s\n", PORT)

	for {
		conn, err := server.Accept()

		if err != nil {
			fmt.Printf("Error accepting: %v\n", err)

			os.Exit(1)
		}

		go handleConnection(conn)
	}
}

func NewTCPServer(port string) (net.Listener) {
	server, err := net.Listen("tcp", ":" + PORT)

	if err != nil {
		fmt.Printf("Error listening to port %s\n", PORT)

		os.Exit(1)
	}

	return server
}

func handleConnection(conn net.Conn) {
}
