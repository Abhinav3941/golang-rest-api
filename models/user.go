package models

import (
	"errors"
	"rest_api_GO/db"
	"rest_api_GO/utlis"
)

type User struct {
	User_id  int64
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

func (u *User) ValidateCredentials() error {

	query := "SELECT User_id, password FROM users WHERE email=?"

	row := db.DB.QueryRow(query, u.Email)

	var retrievedPass string
	row.Scan(&u.User_id, &retrievedPass) // Retrieve User_id and hashed password

	//row := db.DB.QueryRow(query, u.Email)

	//var retrievedPass string // storing password from database with the
	// specific email adreess so that we can compare it to the password from context gin
	//row.Scan(&retrievedPass)

	// but our password is stored in has form so first we need to unhash it and then
	//compare it with the context(user sended) password, but we cannot unhash the stored db password instead
	//we use bcrpyt library comparehashandpassword method which will tell us that user given pass and hashed db password
	//are similar or not

	validPassword := utlis.CheckPasswordHash(u.Password, retrievedPass)
	if validPassword == false {
		return errors.New("credentialsss are invalid")
	}

	return nil
}

func (u *User) Save() error {
	query := `INSERT INTO users (email, password) VALUES (?, ?)` //password should be hased

	statement, err := db.DB.Prepare(query)

	if err != nil {
		panic("put valid user data")
		return err

	}
	defer statement.Close()

	hashedpass, err := utlis.HashPassword(u.Password)
	if err != nil {
		return err
	}

	result, err := statement.Exec(u.Email, hashedpass)

	if err != nil {
		panic("Cant create user Table")

	}

	userId, err := result.LastInsertId()
	u.User_id = userId
	return err

}
