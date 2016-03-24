package main

import (
	"log"
	"net"
)

const PORT = "2525"

func main() {
	server := createServer(PORT)

	defer server.Close()

	for {
		conn, err := server.Accept()

		if err != nil {
			log.Fatalf("Error accepting: %v\n", err)
		}

		handleConnection(conn)
	}
}

func createServer(port string) (net.Listener) {
	server, err := net.Listen("tcp", ":" + PORT)

	if err != nil {
		log.Fatalf("Error listening to port %s\n", PORT)
	}

	return server
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
}
