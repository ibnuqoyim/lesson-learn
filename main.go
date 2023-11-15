package main

import (
	"fmt"
	"net/http"
)

func main() {
	// Menetapkan endpoint dan handler
	http.HandleFunc("/hello", helloHandler)

	// Menjalankan server HTTP di port 8080
	fmt.Println("Server listening on :8080...")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error:", err)
	}
}

// Handler untuk endpoint /hello
func helloHandler(w http.ResponseWriter, r *http.Request) {
	// Menanggapi dengan pesan Hello, World!
	fmt.Fprint(w, "Hello, World!")
}
