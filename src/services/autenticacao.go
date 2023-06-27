package services

import (
	"RestApi/src/models"
	"RestApi/src/utilitarios"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Autenticar(c *gin.Context) {
	var input models.Login

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var usuario models.Usuario
	dbError := models.DB.Where("Email = ?", input.Email).First(&usuario).Error
	if dbError != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Usuário não encontrado"})
		return
	}

	if usuario.Senha != utilitarios.GerarHash(input.Senha) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Senha incorreta"})
		return
	}

	token, err := NewJwtService().GerarToken(usuario.Id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"token": token,
	})
}
