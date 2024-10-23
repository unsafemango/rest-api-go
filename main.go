package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"unsafemango.com/rest-api-go/db"
	"unsafemango.com/rest-api-go/models"
)

func main() {
	db.InitDB()
	server := gin.Default()

	server.GET("/events", getEvents) // handler for incoming http get request

	server.POST("/events", createEvent)

	// starts listening for incoming requests
	server.Run(":8080") // localhost:8080
}

func getEvents(context *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not fetch events. Please try again later.",
		})
		return
	}
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

	err = event.Save()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not create event. Please try again later.",
		})
		return
	}

	context.JSON(http.StatusCreated, gin.H{
		"message": "Event created!",
		"event":   event,
	})
}
