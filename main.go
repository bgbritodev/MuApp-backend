package main

import (
	"log"
	"net/http"

	"github.com/bgbritodev/MuApp-backend/config"
	"github.com/bgbritodev/MuApp-backend/routes"
)

func main() {
	// Conectar ao MongoDB
	config.Connect()

	// Configurar o router
	router := routes.SetupRouter()

	// Iniciar o servidor
	log.Println("Server is running on port 8080")
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatal(err)
	}
}
