package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run . <mode>")
		fmt.Println("Modes: tcp-server, tcp-client, http-server")
		os.Exit(1)
	}

	mode := os.Args[1]
	switch mode {
	case "tcp-server":
		runTCPServer()
	case "tcp-client":
		runTCPClient()
	case "http-server":
		runHTTPServer()
	default:
		fmt.Printf("Unknown mode: %s\n", mode)
		os.Exit(1)
	}
}
