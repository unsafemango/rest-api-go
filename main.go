package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"unsafemango.com/rest-api-go/db"
	"unsafemango.com/rest-api-go/models"
)

func main() {
	db.InitDB()
	server := gin.Default()

	server.GET("/events", getEvents) // handler for incoming http get request

	server.GET("/events/:id", getEvent)

	server.POST("/events", createEvent)

	// starts listening for incoming requests
	server.Run(":8080") // localhost:8080
}

// function to get a list of events
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

// function to get a single event
func getEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Could not parse event id.",
		})
		return
	}

	events, err := models.GetEventByID(eventId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not fetch event.",
		})
		return
	}
	context.JSON(http.StatusOK, events)
}

// function to create an event
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
