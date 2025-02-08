package models

import (
	"rest_api_GO/db"

	"time"
)

type Event struct {
	Evnt_id     int64
	Location    string    `binding:"required"`
	Title       string    `binding:"required"`
	Description string    `binding:"required"`
	DateAndTime time.Time `binding:"required"`
	UserId      int64     `binding:"required"` // who created this event
}

var eventss = []Event{}

func GetEventById(id int64) (*Event, error) {
	query := "SELECT * FROM eventss WHERE Evnt_id = ? "

	row := db.DB.QueryRow(query, id)

	var event Event

	err := row.Scan(
		&event.Evnt_id,
		&event.Location,
		&event.Title,
		&event.DateAndTime, // Correct order
		&event.Description, // Now at the right index
		&event.UserId)
	if err != nil {
		return nil, err
	}

	return &event, nil

}

// creating in eventhandler and save event
func (e *Event) Save() error {
	// later add to DB

	query := `INSERT INTO eventss("location", "title", "description", "dateandtime", "user_id" )

VALUES ( ? ,? ,? ,?,?)`

	statement, err := db.DB.Prepare(query)
	if err != nil {
		panic(err)
	}

	defer statement.Close()
	result, err := statement.Exec(e.Location, e.Title, e.Description, e.DateAndTime, e.UserId)
	if err != nil {
		panic(err)
	}

	E_id, err := result.LastInsertId() // to get the event id bcz its auto incremented
	e.Evnt_id = E_id
	return err

}

func GetAllEvents() ([]Event, error) {

	query := "SELECT * FROM eventss"

	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}

	var evnts []Event // slice to store all the data from the query

	for rows.Next() {

		var evnt Event

		err := rows.Scan(
			&evnt.Evnt_id,
			&evnt.Location,
			&evnt.Title,
			&evnt.DateAndTime, // Correct order
			&evnt.Description, // Now at the right index
			&evnt.UserId,
		)
		if err != nil {
			return nil, err
		}
		evnts = append(evnts, evnt)

	}
	return evnts, nil
}

// it is  method
func (event Event) UpdateEvent() error {
	query := `
UPDATE eventss
SET location = ? ,title=?,description=?,dateandtime=?
WHERE Evnt_id = ?
`
	statement, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}

	defer statement.Close()

	_, err = statement.Exec(
		event.Location,
		event.Title,
		event.Description,
		event.DateAndTime,
		event.Evnt_id)
	return err
}

func (event Event) DeleteEventById() error {

	query := "DELETE FROM eventss WHERE Evnt_id = ?"

	statement, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer statement.Close()

	_, err = statement.Exec(event.Evnt_id)
	return err

}

func (e Event) Register(userid int64) error {

	query := "INSERT INTO registrations(E_id , U_id) VALUES (?,?)"
	statement, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer statement.Close()
	_, err = statement.Exec(e.UserId, userid)

	return err

}

func (e Event) CancelRegistration(userid int64) error {
	query := "DELETE FROM registrations WHERE E_id = ? AND U_id = ?"
	statement, err := db.DB.Prepare(query)
	if err != nil {
		return err

	}
	defer statement.Close()
	_, err = statement.Exec(e.UserId, userid)

	return err
}
