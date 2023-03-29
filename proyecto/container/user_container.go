package container

import (
	"proyecto/controllers"
	"proyecto/repositories"
	"proyecto/services"
)

type UserContainer struct {
	UserController controllers.UserController
	UserService    services.UserService
	UserRepository repositories.UserRepository
}

func NewUserContainer() *controllers.UserController {
	userRepo := repositories.NewUserRepository()
	userService := services.NewUserService(userRepo)
	userController := controllers.NewUserController(userService)

	return &userController
}
