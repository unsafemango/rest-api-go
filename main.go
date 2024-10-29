package main

import (
	"github.com/gin-gonic/gin"
	"unsafemango.com/rest-api-go/db"
	"unsafemango.com/rest-api-go/routes"
)

func main() {
	db.InitDB()
	server := gin.Default()

	routes.RegisterRoutes(server)

	// starts listening for incoming requests
	server.Run(":8080") // localhost:8080
}
