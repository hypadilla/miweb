package container

import (
	"proyecto/controllers"
	"proyecto/repositories"
	"proyecto/services"
)

type Container struct {
	userController *controllers.UserController
	productController *controllers.ProductController
}

func NewContainer() *Container {
	userRepository := repositories.NewUserRepository()
	userService := services.NewUserService(userRepository)
	userController := controllers.NewUserController(userService)

	productRepository := repositories.NewProductRepository()
	productService := services.NewProductService(productRepository)
	productController := controllers.NewProductController(productService)

	return &Container{
		userController: &userController,
		productController: &productController,
	}
}
