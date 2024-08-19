package routes

import (
	"github.com/bgbritodev/MuApp-backend/controllers"
	"github.com/gin-gonic/gin"
)

// SetupUserRoutes configura as rotas para o controlador de usuários usando Gin
func SetupUserRoutes(router *gin.Engine) {
	// Rota para criação de usuários
	router.POST("/users/create", controllers.CreateUser)

	// Rota para edição de usuários
	router.PUT("/users/edit/:id", controllers.UpdateUser)

	// Rota para login de usuários
	router.POST("/users/login", controllers.LoginUser)
}
