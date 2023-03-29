package http

import (
	"proyecto/controllers"
	
	"github.com/gin-gonic/gin"
)

func NewHTTPRouter(userController controllers.UserController, 
    productController controllers.ProductController, 
    authController controllers.AuthController) *gin.Engine {
    r := gin.Default()
    
    userRoutes := r.Group("/api")
    RegisterUserRoutes(userRoutes, userController)
    
   	productRoutes := r.Group("/api")
    RegisterProductRoutes(productRoutes, productController)

    authRoutes := r.Group("/api")
    //RegisterAuthRoutes(authRoutes, authController)
    NewAuthRouter(authRoutes, authController)

    return r
}