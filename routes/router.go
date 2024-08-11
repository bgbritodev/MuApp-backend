package routes

import (
	"github.com/gorilla/mux"
)

func SetupRouter() *mux.Router {
	router := mux.NewRouter()

	// Configurar rotas de obras
	SetupObraRoutes(router)

	SetupMuseuRoutes(router)

	SetupSalaRoutes(router)

	return router
}
