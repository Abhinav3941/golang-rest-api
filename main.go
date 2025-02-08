package main

import (
	"github.com/gin-gonic/gin"
	"rest_api_GO/db"
	"rest_api_GO/routes"
)

func main() {
	db.Initdb()
	server := gin.Default()
	routes.RegisterRoutes(server)
	server.Run(":8080")

}
