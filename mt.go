package main

import (
	"fmt"
	"log"
	"net"
	"os"
)

const PORT = "2525"

func main() {
	server := CreateServer(PORT)

	defer server.Close()

	for {
		conn, err := server.Accept()

		if err != nil {
			fmt.Printf("Error accepting: %v\n", err)

			os.Exit(1)
		}

		go handleConnection(conn)
	}
}

func CreateServer(port string) (net.Listener) {
	server, err := net.Listen("tcp", ":" + PORT)

	if err != nil {
		log.Fatalf("Error listening to port %s\n", PORT)
	}

	return server
}

func handleConnection(conn net.Conn) {
}
