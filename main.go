package main

import (
	db "example.com/rest-api/DB"
	"example.com/rest-api/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()
	db.InitDB()

	routes.RegisterRoutes((server))
	server.Run(":8080")
}
