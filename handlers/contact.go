package handlers

import (
	"fmt"
	"log"
	"net/http"

	mailgun "github.com/mailgun/mailgun-go"
)

//Contact handles the contact form
func Contact(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	email := r.FormValue("email")
	msg := r.FormValue("msg")
	fmt.Println(name, email, msg)

	mg := mailgun.NewMailgun("sandboxd6d9d06b590b44d7b8dcacae8c7e1528.mailgun.org", "key-27a0d63a966e0691e65677ccf99bf804", "https://api.mailgun.net/v3/sandboxd6d9d06b590b44d7b8dcacae8c7e1528.mailgun.org")
	mail := mailgun.NewMessage(
		"Excited User <"+email+">",
		"Test Subject",
		msg,
		"huzi8014@gmail.com",
		"huzaifaa.2002@gmail.com",
	)
	res, id, err := mg.Send(mail)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Message id=%s, Message Res:%s", id, res)
}
