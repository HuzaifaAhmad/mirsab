package models

import (
	"errors"
	"log"
	"time"
)

//Contact stores contact info
type Contact struct {
	ID      int
	Name    string
	Email   string
	Subject string
	Msg     string
	Time    string
}

//AddContact => contact table in DB
func AddContact(cntc Contact) error {
	stmt := `INSERT INTO contact (name, email, subject, msg, created_at) VALUES ($1, $2, $3, $4, $5) RETURNING id;`
	id := 0
	err := DB.QueryRow(stmt, cntc.Name, cntc.Email, cntc.Subject, cntc.Msg, cntc.Time).Scan(&id)
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
	stmt := `SELECT id, name, email, subject, msg, created_at FROM contact;`
	rows, err := DB.Query(stmt)
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()
	for rows.Next() {
		var c Contact
		err = rows.Scan(&c.ID, &c.Name, &c.Email, &c.Subject, &c.Msg, &c.Time)
		if err != nil {
			log.Fatal(err)
			return []Contact{}, nil
		}
		ti, err := time.Parse("2006-01-02T15:04:00Z", c.Time)
		if err != nil {
			log.Fatal(err)
		}
		c.Time = ti.Format("Jan 2, 2006 3:04PM")

		cont = append(cont, c)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
		return []Contact{}, nil
	}

	return cont, nil
}
