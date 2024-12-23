package main

import (
	"fmt"
	"net/http"
)

func runHTTPServer() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("Got %s request to %s\n", r.Method, r.URL.Path)
		fmt.Fprintf(w, "Hello from Go!")
	})

	fmt.Println("HTTP server listening on :8080")
	http.ListenAndServe(":8080", nil)
}
