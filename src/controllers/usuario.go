package controllers

import (
	"RestApi/src/services"

	"github.com/gin-gonic/gin"
)

func Cadastrar(c *gin.Context) {
	services.UsuarioCadastrar(c)
}

func ListarUsuarios(c *gin.Context) {
	services.UsuarioListar(c)
}
