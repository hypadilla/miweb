package container

import (
	"proyecto/controllers"
	"proyecto/repositories"
	"proyecto/services"
)

type ProductContainer struct {
	ProductController controllers.ProductController
	ProductService    services.ProductService
	ProductRepository repositories.ProductRepository
}

func NewProductContainer() *controllers.ProductController {
	productRepo := repositories.NewProductRepository()
	productService := services.NewProductService(productRepo)
	productController := controllers.NewProductController(productService)

	return &productController
}
