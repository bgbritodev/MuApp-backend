package routes

import (
	"github.com/bgbritodev/MuApp-backend/controllers"
	"github.com/gorilla/mux"
)

// SetupSalaRoutes configura as rotas para o controlador de salas
func SetupSalaRoutes(router *mux.Router) {
	router.HandleFunc("/salas", controllers.CreateSala).Methods("POST")
	router.HandleFunc("/salas/{id}", controllers.GetSala).Methods("GET")
	router.HandleFunc("/salas/museu/{museuId}", controllers.GetSalasByMuseuID).Methods("GET")
}
