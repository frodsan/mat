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

		go handleConnection(conn)
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

	conn.Write([]byte("220 OK\n"))

	for i := 1; i <= 5; i++ {
		buffer := make([]byte, 1024)

		n, _ := conn.Read(buffer)

		if string(buffer[:n]) == "DATA\r\n" {
			break
		}

		conn.Write([]byte("250 OK\n"))
	}
}
