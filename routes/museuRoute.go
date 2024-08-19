package routes

import (
	"github.com/bgbritodev/MuApp-backend/controllers"
	"github.com/gin-gonic/gin"
)

func SetupMuseuRoutes(router *gin.Engine) {
	// Grupo de rotas para museus
	museuRoutes := router.Group("/museus")
	{
		museuRoutes.POST("/", controllers.CreateMuseu)
		museuRoutes.GET("/:id", controllers.GetMuseu)
		museuRoutes.GET("/all", controllers.GetAllMuseus)
		museuRoutes.PUT("/:id", controllers.UpdateMuseu)
		museuRoutes.DELETE("/:id", controllers.DeleteMuseu)
	}
}
