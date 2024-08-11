package routes

import (
	"github.com/bgbritodev/MuApp-backend/controllers"
	"github.com/gorilla/mux"
)

// SetupObraRoutes configura as rotas para o controlador de obras
func SetupObraRoutes(router *mux.Router) {
	router.HandleFunc("/obras", controllers.CreateObras).Methods("POST")
	router.HandleFunc("/obras/{id}", controllers.GetObra).Methods("GET")
}
