package models

import (
	"errors"
	"log"
)

//Contact stores contact info
type Contact struct {
	Name    string
	Email   string
	Subject string
	Msg     string
}

//AddContact => contact table in DB
func AddContact(cntc Contact) error {
	stmt := `INSERT INTO contact (name, email, subject, msg) VALUES ($1, $2, $3, $4) RETURNING id;`
	id := 0
	err := DB.QueryRow(stmt, cntc.Name, cntc.Email, cntc.Subject, cntc.Msg).Scan(&id)
	if err != nil {
		log.Fatal(err)
		err := errors.New("could not process contact form")
		return err
	}
	return nil
}

//GetContacts gets information from DB
func GetContacts() ([]Contact, error) {
	var cont []Contact
	stmt := `SELECT name, email, subject, msg FROM contact;`
	rows, err := DB.Query(stmt)
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()
	for rows.Next() {
		var c Contact
		err = rows.Scan(&c.Name, &c.Email, &c.Subject, &c.Msg)
		if err != nil {
			log.Fatal(err)
			return []Contact{}, nil
		}
		cont = append(cont, c)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
		return []Contact{}, nil
	}

	return cont, nil
}
