package models

import (
	"database/sql"
	"errors"
	"fmt"
)

type User struct {
	ID        int
	FirstName string
	LastName  string
	Email     string
	Password  string
}

//AddUser adds user returning an error if occured
func AddUser(fname, lname, email, pass string) error {
	stmt := `INSERT INTO users (fname, lname, email, password) VALUES ($1, $2, $3, $4) RETURNING id;`
	id := 0
	err := DB.QueryRow(stmt, fname, lname, email, pass).Scan(&id)
	if err != nil {
		fmt.Println(err)
		newerr := errors.New("could not add user")
		return newerr
	}
	return nil
}

//GetUser adds user returning an error if occured
func GetUser(eml, pass string) (User, error) {
	var u User
	stmt := `SELECT * FROM users where email=$1;`
	row := DB.QueryRow(stmt, eml)
	err := row.Scan(&u.ID, &u.FirstName, &u.LastName, &u.Email, &u.Password)
	switch err {
	case sql.ErrNoRows:
		return User{}, sql.ErrNoRows
	case nil:
		return u, nil
	default:
		panic(err)
	}
}
