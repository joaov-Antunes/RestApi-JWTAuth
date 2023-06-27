package middlewares

import (
	"RestApi/src/services"

	"github.com/gin-gonic/gin"
)

func Autenticacao() gin.HandlerFunc {
	return func(c *gin.Context) {
		const Bearer = "Bearer "
		header := c.GetHeader("Authorization")
		if header == "" {
			c.AbortWithStatus(401)
		}

		token := header[len(Bearer):]

		if !services.NewJwtService().ValidarToken(token) {
			c.AbortWithStatus(401)
		}
	}
}
