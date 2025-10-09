package main

import (
	"fmt"
	"net/http"
)

// Endpoint básico: http://localhost:8080/api/hello
func main() {
	http.HandleFunc("/api/hello", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, `{"message": "Hello from Go backend!"}`)
	})

	fmt.Println("✅ Go server running on port 8080...")
	http.ListenAndServe(":8080", nil)
}
