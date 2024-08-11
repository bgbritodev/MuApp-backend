package routes

import (
	"github.com/bgbritodev/MuApp-backend/controllers"
	"github.com/gorilla/mux"
)

func SetupRoutes(router *mux.Router) {
	router.HandleFunc("/obras", controllers.CreateObra).Methods("POST")
	router.HandleFunc("/obras/{id}", controllers.GetObra).Methods("GET")
}
