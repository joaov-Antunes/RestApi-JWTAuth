package controllers

import (
	"RestApi/src/services"

	"github.com/gin-gonic/gin"
)

func Autenticar(c *gin.Context) {
	services.Autenticar(c)
}
