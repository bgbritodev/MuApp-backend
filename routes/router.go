package routes

import (
	"github.com/gin-gonic/gin"
)

// SetupRouter configura as rotas da aplicação usando Gin
func SetupRouter() *gin.Engine {
	router := gin.Default()

	// Configurar rotas de obras
	SetupObraRoutes(router)

	// Configurar rotas de museus
	SetupMuseuRoutes(router)

	// Configurar rotas de salas
	SetupSalaRoutes(router)

	// Configurar rotas de usuários
	SetupUserRoutes(router)

	return router
}
