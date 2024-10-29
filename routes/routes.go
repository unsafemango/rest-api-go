package routes

import "github.com/gin-gonic/gin"

func RegisterRoutes(server *gin.Engine) {
	server.GET("/events", getEvents) // handler for incoming http get request

	server.GET("/events/:id", getEvent)

	server.POST("/events", createEvent)
}
