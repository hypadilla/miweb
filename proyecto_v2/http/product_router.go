package http

import (
	"proyecto/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterProductRoutes(router *gin.RouterGroup, productController controllers.ProductController) {
	router.GET("/products", productController.GetAllProducts)
	router.GET("/products/:id", productController.GetProductByID)
	router.POST("/products", productController.CreateProduct)
	router.PUT("/products/:id", productController.UpdateProduct)
	router.DELETE("/products/:id", productController.DeleteProduct)
}

func NewProductRouter(productController controllers.ProductController) *gin.RouterGroup {
	r := gin.New()
	productRouter := r.Group("/api")
	RegisterProductRoutes(productRouter, productController)
	return productRouter
}
