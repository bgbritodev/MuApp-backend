package main

import (
	"fmt"
	"net/http"

	"github.com/bgbritodev/MuApp-backend/config"
)

func main() {
	config.Connect()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, MuApp-backend!")
	})

	fmt.Println("Server started at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
