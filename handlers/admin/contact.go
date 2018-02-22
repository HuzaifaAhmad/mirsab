package admin

import (
	"log"
	"net/http"

	"github.com/HuzaifaAhmad/mirsab/models"
)

var c models.Contact

//Contact renders the admin contact page
func Contact(w http.ResponseWriter, r *http.Request) {
	c, err := models.GetContacts()
	if err != nil {
		log.Fatal(err)
	}

	email := getSession(r)
	if email != "" {
		err := tpl.ExecuteTemplate(w, "contact.gohtml", c)
		if err != nil {
			log.Fatalf("template execution: %s", err)
		}
	} else {
		http.Redirect(w, r, "/login", 302)
	}
}
