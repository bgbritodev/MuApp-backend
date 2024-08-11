package routes

import (
	"github.com/bgbritodev/MuApp-backend/controllers"
	"github.com/gorilla/mux"
)

func RegisterRoutes(r *mux.Router) {
	r.HandleFunc("/upload", controllers.UploadMedia).Methods("POST")
	r.HandleFunc("/media/{id}", controllers.GetMedia).Methods("GET")
}
