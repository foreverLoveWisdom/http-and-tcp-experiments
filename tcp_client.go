package main

import (
	"fmt"
	"net"
)

func runTCPClient() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Printf("Failed to connect: %v\n", err)
		return
	}
	defer conn.Close()

	message := "Hello, server!"
	_, err = conn.Write([]byte(message))
	if err != nil {
		fmt.Printf("Failed to send: %v\n", err)
		return
	}

	buffer := make([]byte, 1024)
	n, err := conn.Read(buffer)
	if err != nil {
		fmt.Printf("Failed to read response: %v\n", err)
		return
	}

	fmt.Printf("Server says: %s\n", string(buffer[:n]))
}
