package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"unsafemango.com/rest-api-go/models"
)

func main() {
	server := gin.Default()

	server.GET("/events", getEvents) // handler for incoming http get request

	server.POST("/events", createEvent)

	// starts listening for incoming requests
	server.Run(":8080") // localhost:8080
}

func getEvents(context *gin.Context) {
	events := models.GetAllEvents()
	context.JSON(http.StatusOK, events)
}

func createEvent(context *gin.Context) {
	var event models.Event
	err := context.ShouldBindJSON(&event)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Could not parse request data",
		})
		return
	}

	event.ID = 1
	event.UserID = 1

	context.JSON(http.StatusCreated, gin.H{
		"message": "Event created!",
		"event":   event,
	})
}