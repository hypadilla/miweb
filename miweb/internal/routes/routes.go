package routes

import (
	"net/http"
	"paquetes/internal/handlers"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.Default()

	// Ruta de ejemplo
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hola mundo!",
		})
	})

	// Ruta de ejemplo
	r.GET("/user", handlers.Create)

	return r
}
