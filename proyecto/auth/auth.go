package auth

import (
	"fmt"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func Authenticate() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Se requiere un token de autenticación"})
			return
		}

		// Obtener el token de la cabecera "Authorization" en el formato "Bearer <token>"
		tokenString := authHeader[len("Bearer "):]

		// Validar el token
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			// Verifica que el token fue firmado con la clave correcta
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("método de firma inesperado: %v", token.Header["alg"])
			}

			// Retorna la clave secreta para verificar la firma del token
			return []byte("secret"), nil
		})

		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Token de acceso inválido" + err.Error()})
			return
		}

		if token.Valid {
			/* //Verifica que el usuario tenga los permisos necesarios para acceder al recurso solicitado
			if permiso := claims["permiso"]; permiso != "admin" {
				c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "No tienes permiso para acceder a este recurso"})
				return
			}*/

			// Guardar la identidad del usuario en el contexto de Gin
			//c.Set("user_id", claims["id"].(string))

			// Llama al siguiente handler si el token es válido y el usuario tiene los permisos necesarios
			c.Next()
		} else {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Token de acceso inválido"})
			return
		}
	}
}
