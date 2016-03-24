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
		go handleConnection(acceptConnection(server))
	}
}

func createServer(port string) (net.Listener) {
	server, err := net.Listen("tcp", ":" + PORT)

	if err != nil {
		log.Fatalf("Error listening to port %s\n", PORT)
	}

	return server
}

func acceptConnection(server net.Listener) net.Conn {
	conn, err := server.Accept()

	if err != nil {
		log.Fatalf("Error accepting: %v\n", err)
	}

	return conn
}

func handleConnection(conn net.Conn) {
}
