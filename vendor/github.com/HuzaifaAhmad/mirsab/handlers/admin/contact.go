package admin

import (
	"log"
	"net/http"
	"strconv"

	"github.com/HuzaifaAhmad/mirsab/models"
	"github.com/gorilla/mux"
)

//Contact renders the admin contact page
func Contact(w http.ResponseWriter, r *http.Request) {
	c, err := models.GetContacts()
	if err != nil {
		log.Fatal(err)
	}
	err = tpl.ExecuteTemplate(w, "contact.gohtml", c)
	if err != nil {
		log.Fatalf("template execution: %s", err)
	}
}

//ContactDetails handles details for contacts on admin panel
func ContactDetails(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id := vars["id"]
	intg, err := strconv.Atoi(id)
	if err != nil {
		log.Fatal(err)
	}
	c, err := models.GetContacts()
	if err != nil {
		log.Fatal(err)
	}
	var cnt models.Contact
	//checking db against url param
	for _, val := range c {
		if val.ID == intg {
			cnt = val
			break
		}
	}

	err = tpl.ExecuteTemplate(w, "contactDetails.gohtml", cnt)
	if err != nil {
		log.Fatalf("template execution: %s", err)
	}

}
