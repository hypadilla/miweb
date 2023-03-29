package http

import (
	"proyecto/auth"
	"proyecto/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(router *gin.RouterGroup, userController controllers.UserController) {
	router.GET("/users", auth.Authenticate(), userController.GetAllUsers)
	router.GET("/users/:id", auth.Authenticate(), userController.GetUserByID)
	router.POST("/users", auth.Authenticate(), userController.CreateUser)
	router.PUT("/users/:id", auth.Authenticate(), userController.UpdateUser)
	router.DELETE("/users/:id", auth.Authenticate(), userController.DeleteUser)
}

func NewUserRouter(userController controllers.UserController) *gin.RouterGroup {
	r := gin.New()
	userRouter := r.Group("/api")
	RegisterUserRoutes(userRouter, userController)
	return userRouter
}
