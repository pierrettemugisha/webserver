package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

func handleRequest(conn net.Conn) {
	defer conn.Close()

	reader := bufio.NewReader(conn)

	requestLine, err := reader.ReadString('\n')

	if err != nil {
		// handle error
		fmt.Println("Error reading request: ", err)
		return
	}

	parts := strings.Fields(requestLine)

	if len(parts) < 2 {
		fmt.Println("Invalid Request: ", requestLine)
		return
	}

	path := parts[1]
	response := fmt.Sprintf("HTTP/1.1 200 OK\r\n\r\nRequested path: %s\r\n", path)

	conn.Write([]byte(response))
}

func main() {
	// fmt.Println("Hello, World!")

	listener, err := net.Listen("tcp", "127.0.0.1:80")

	if err != nil {
		// handle error
		fmt.Println("Error creating listener: ", err)
		return
	}

	defer listener.Close()

	fmt.Println("Server is listening on port 80...")

	// Accept and handle incoming connections

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error creating listener: ", err)
			continue
		}
		go handleRequest(conn)
	}
}
