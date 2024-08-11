package main

import (
	"fmt"
	"net/http"

	"github.com/bgbritodev/MuApp-backend/config"
	"github.com/bgbritodev/MuApp-backend/routes"
	"github.com/gorilla/mux"
)

func main() {
	config.Connect()

	r := mux.NewRouter()
	routes.RegisterRoutes(r)

	http.Handle("/", r)

	fmt.Println("Server started at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
