package models

import "time"

type Event struct {
	ID          int
	Name        string    `binding:"required"`
	Description string    `binding:"required"`
	Location    string    `binding:"required"`
	DateTime    time.Time `binding:"required"`
	UserID      int       `binding:"required"`
}

// func (e Event) Save() {
// 	 query := "INSERT INTO events(name,de)"

// }
