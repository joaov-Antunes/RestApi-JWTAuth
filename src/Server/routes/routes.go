package routes

import (
	"RestApi/src/Server/middlewares"
	"RestApi/src/controllers"

	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine) *gin.Engine {
	main := router.Group("/v1")
	{
		usuario := main.Group("/user", middlewares.Autenticacao())
		{
			usuario.GET("/usuario-listar", controllers.ListarUsuarios)
		}

		main.POST("/usuario-cadastrar", controllers.Cadastrar)
		main.POST("/autenticar", controllers.Autenticar)
	}

	return router
}
