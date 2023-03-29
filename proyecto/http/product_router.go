package http

import (
	"proyecto/auth"
	"proyecto/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterProductRoutes(router *gin.RouterGroup, productController controllers.ProductController) {
	router.GET("/products", auth.Authenticate(), productController.GetAllProducts)
	router.GET("/products/:id", auth.Authenticate(), productController.GetProductByID)
	router.POST("/products", auth.Authenticate(), productController.CreateProduct)
	router.PUT("/products/:id", auth.Authenticate(), productController.UpdateProduct)
	router.DELETE("/products/:id", auth.Authenticate(), productController.DeleteProduct)
}

func NewProductRouter(productController controllers.ProductController) *gin.RouterGroup {
	r := gin.New()
	productRouter := r.Group("/api")
	RegisterProductRoutes(productRouter, productController)
	return productRouter
}
