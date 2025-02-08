package db

import (
	"database/sql"
	_ "modernc.org/sqlite"
)

var DB *sql.DB

func Initdb() {

	var err error

	DB, err = sql.Open("sqlite", "./mydb.db")

	if err != nil {
		panic("cannot connect to db" + err.Error())
	}

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)
	CreateTable()
}

func CreateTable() {

	UserTable :=
		`CREATE TABLE IF NOT EXISTS users (
    User_id INTEGER PRIMARY KEY AUTOINCREMENT,
    Email TEXT NOT NULL UNIQUE,
    Password TEXT NOT NULL
)`
	_, err := DB.Exec(UserTable)

	if err != nil {
		panic("could not create users table" + err.Error())
	}

	CreateEventTable :=
		`CREATE TABLE IF NOT EXISTS eventss (
    EVNT_ID INTEGER PRIMARY KEY AUTOINCREMENT,
    LOCATION TEXT NOT NULL,
    TITLE TEXT NOT NULL,
    DATEANDTIME DATETIME NOT NULL,
    DESCRIPTION TEXT NOT NULL,
    USER_ID INTEGER NOT NULL,
    FOREIGN KEY (USER_ID) REFERENCES users(User_id)
)`

	_, err = DB.Exec(CreateEventTable)

	if err != nil {
		panic("cannot create event table" + err.Error())
	}

	CreateRegistrationTable :=
		`
 CREATE TABLE  IF NOT EXISTS registrations (
R_id INTEGER PRIMARY KEY AUTOINCREMENT,
E_id INTEGER NOT NULL,
U_id INTEGER NOT NULL,
FOREIGN KEY(E_id) REFERENCES eventss(EVNT_ID),
FOREIGN KEY(U_ID) REFERENCES users(User_id)
)
`

	_, err = DB.Exec(CreateRegistrationTable)
	if err != nil {
		panic("cannot create registration table" + err.Error())
	}
}
