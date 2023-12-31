package server

import (
	"RestApi/src/Server/routes"
	"log"

	"github.com/gin-gonic/gin"
)

type Server struct {
	port   string
	server *gin.Engine
}

func NewServer() Server {
	return Server{
		port:   "3000",
		server: gin.Default(),
	}
}

func (s *Server) Run() {
	router := routes.Routes(s.server)

	log.Printf("Server running at port: %v", s.port)
	log.Fatal(router.Run(":" + s.port))
}
