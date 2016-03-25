package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

const PORT = "2525"

func main() {
	server := createServer(PORT)

	fmt.Println("Listening on port", PORT)

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

	buffer := bufio.NewReader(conn)

	for {
		str, _ := buffer.ReadString('\n')

		if str == "DATA\r\n" {
			break
		}

		conn.Write([]byte("250 OK\n"))
	}

	conn.Write([]byte("354 OK\n"))

	buffer = bufio.NewReader(conn)

	fmt.Println("---")

	for {
		str, _ := buffer.ReadString('\n')

		if str == ".\r\n" {
			break
		}

		fmt.Print(str)
	}

	fmt.Println("")

	conn.Write([]byte("250 OK\n"))
	conn.Write([]byte("221 OK\n"))
}
