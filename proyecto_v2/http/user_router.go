package http

import (
	"proyecto/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(router *gin.RouterGroup, userController controllers.UserController) {
	router.GET("/users", userController.GetAllUsers)
	router.GET("/users/:id", userController.GetUserByID)
	router.POST("/users", userController.CreateUser)
	router.PUT("/users/:id", userController.UpdateUser)
	router.DELETE("/users/:id", userController.DeleteUser)
}

func NewUserRouter(userController controllers.UserController) *gin.RouterGroup {
	r := gin.New()
	userRouter := r.Group("/api")
	RegisterUserRoutes(userRouter, userController)
	return userRouter
}
