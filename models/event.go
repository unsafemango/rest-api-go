package models

import "time"

type Event struct {
	ID          int
	Name        string    `binding:"required"`
	Description string    `binding:"required"`
	Location    string    `binding:"required"`
	DateTime    time.Time `binding:"required"`
	UserID      int
}

var events = []Event{}

// method to save an event
func (e Event) Save() {
	// later add to database
	events = append(events, e)
}

// function to get all events -  call it to call all available events not on an existing event
func GetAllEvents() []Event {
	return events
}
