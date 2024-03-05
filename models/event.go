package models

import (
	"time"

	"github.com/SureshkumarUndala/Event_Management_API/db"
)

type Event struct {
	Id          int       `json:"id"`
	Name        string    `json:"name" binding:"required"`
	Description string    `json:"description" binding:"required"`
	Location    string    `json:"location" binding:"required"`
	DateTime    time.Time `json:"datetime"`
	UserID      int       `json:"userid"`
}

func (e Event) Save() (int64, error) {

	query := "INSERT INTO events(name,description,location,dateTime,user_id) VALUES(?,?,?,?,?)"

	rows, err := db.DB.Exec(query, e.Name, e.Description, e.Location, e.DateTime, e.UserID)

	if err != nil {
		return 0, err
	}
	id, err := rows.LastInsertId()
	if err != nil {
		return 0, err
	}
	return id, nil

}

func GetallEvents() ([]Event, error) {
	var events []Event

	query := "SELECT  eventId, name, description,location,user_id from events"

	rows, err := db.DB.Query(query)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var event Event
		err := rows.Scan(&event.Id, &event.Name, &event.Description, &event.Location, &event.UserID)
		if err != nil {
			return nil, err
		}

		events = append(events, event)
	}
	return events, nil

}

func GetEvent(id int) (Event, error) {
	var event Event
	query := "SELECT * FROM events where eventid = ?"
	err := db.DB.QueryRow(query, id).Scan(&event)
	if err != nil {
		return event, err
	}
	return event, nil
}
