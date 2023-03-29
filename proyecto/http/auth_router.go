package http

import (
	"proyecto/controllers"

	"github.com/gin-gonic/gin"
)

func registerAuthRoutes(router *gin.RouterGroup, authController controllers.AuthController) {
	router.POST("/login", authController.Login)
	router.POST("/register", authController.Register)
}

func NewAuthRouter(router *gin.RouterGroup, authController controllers.AuthController){
	registerAuthRoutes(router, authController)
}
