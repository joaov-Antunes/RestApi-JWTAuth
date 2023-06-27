package services

import (
	"RestApi/src/models"
	"RestApi/src/utilitarios"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UsuarioCadastrar(c *gin.Context) {
	var input models.Usuario
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	senha := utilitarios.GerarHash(input.Senha)

	usuario := models.Usuario{Nome: input.Nome, Email: input.Email, Senha: senha}
	models.DB.Create(&usuario)

	c.JSON(http.StatusOK, gin.H{"Data": usuario})
}

func UsuarioListar(c *gin.Context) {
	var usuario []models.Usuario
	models.DB.Where("email = ?", "joaovicenteftn@gmail.com").Find(&usuario)

	c.JSON(http.StatusOK, gin.H{"Data": usuario})
}
