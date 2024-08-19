package routes

import (
	"github.com/bgbritodev/MuApp-backend/controllers"
	"github.com/gin-gonic/gin"
)

// SetupSalaRoutes configura as rotas para o controlador de salas
func SetupSalaRoutes(router *gin.Engine) {
	// Rota para criar uma nova sala
	router.POST("/salas", controllers.CreateSala)

	// Rota para recuperar uma sala específica pelo ID
	router.GET("/salas/:id", controllers.GetSala)

	// Rota para recuperar todas as salas de um determinado museu
	router.GET("/salas/museu/:museuId", controllers.GetSalasByMuseuID)
}
