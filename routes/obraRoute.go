package routes

import (
	"github.com/bgbritodev/MuApp-backend/controllers"
	"github.com/gin-gonic/gin"
)

// SetupObraRoutes configura as rotas para o controlador de obras usando Gin
func SetupObraRoutes(router *gin.Engine) {
	router.POST("/obras", controllers.CreateObra)
	router.GET("/obras/:id", controllers.GetObra)
	router.GET("/obras/sala/:salaId", controllers.GetObrasBySalaID)
}
