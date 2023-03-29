package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"paquetes/miweb/internal/services"
)

type UsersHandler struct {
	usersService *services.UsersService
}

func NewUsersHandler(usersService *services.UsersService) *UsersHandler {
	return &UsersHandler{
		usersService: usersService,
	}
}

func (h *UsersHandler) Create(c *gin.Context) {
	// Parsear los datos del usuario que vienen en el cuerpo de la solicitud
	var user services.User
	err := c.BindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Llamar al servicio de usuarios para crear el nuevo usuario
	err = h.usersService.Create(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Devolver la respuesta con el nuevo usuario creado
	c.JSON(http.StatusCreated, user)
}