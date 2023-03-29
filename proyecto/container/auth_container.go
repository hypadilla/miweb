package container

import (
	"proyecto/controllers"
	"proyecto/repositories"
	"proyecto/services"
)

type AuthContainer struct {
	authService services.AuthService
	userRepo    repositories.UserRepository
	authCtrl    controllers.AuthController
}

func NewAuthContainer() *controllers.AuthController {
	userRepo := repositories.NewUserRepository()
	authService := services.NewAuthService(userRepo, "secret", 86400)
	authCtrl := controllers.NewAuthController(authService)

	return &authCtrl
}
