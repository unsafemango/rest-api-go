package models

import "time"

type Event struct {
	ID          int
	Name        string
	Description string
	Location    string
	DateTime    time.Time
	UserID      int
}

var events = []Event{}

// method to save an event
func (e Event) Save() {
	// later add to database
	events = append(events, e)
}
