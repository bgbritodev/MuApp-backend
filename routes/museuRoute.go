package routes

import (
	"github.com/bgbritodev/MuApp-backend/controllers"
	"github.com/gorilla/mux"
)

func SetupMuseuRoutes(router *mux.Router) {
	router.HandleFunc("/museus", controllers.CreateMuseu).Methods("POST")
	router.HandleFunc("/museus/{id}", controllers.GetMuseu).Methods("GET")
	router.HandleFunc("/allmuseus", controllers.GetAllMuseus).Methods("GET")
	router.HandleFunc("/museus/{id}", controllers.UpdateMuseu).Methods("PUT")
	router.HandleFunc("/museus/{id}", controllers.DeleteMuseu).Methods("DELETE")
}
