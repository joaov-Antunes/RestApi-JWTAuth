package main

import (
	server "RestApi/src/Server"
	"RestApi/src/models"
)

func main() {
	models.ConnectDatabase()

	server := server.NewServer()

	server.Run()
}
