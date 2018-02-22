package handlers

import (
	"html/template"
	"log"
	"net/http"

	"github.com/HuzaifaAhmad/mirsab/models"
)

type flashMsg struct {
	Msg string
}

var flash flashMsg

var cntc models.Contact

//ContactPage renders the contact template
func ContactPage(w http.ResponseWriter, r *http.Request, tpl *template.Template) {
	err := tpl.ExecuteTemplate(w, "contact.gohtml", flash)
	if err != nil {
		log.Fatalf("template execution: %s", err)
	}

	flash.Msg = ""
}

//Contact handles the contact form
func Contact(w http.ResponseWriter, r *http.Request) {
	cntc.Name = r.FormValue("name")
	cntc.Email = r.FormValue("email")
	cntc.Subject = r.FormValue("subject")
	cntc.Msg = r.FormValue("msg")

	// mg := mailgun.NewMailgun("sandboxd6d9d06b590b44d7b8dcacae8c7e1528.mailgun.org", "key-27a0d63a966e0691e65677ccf99bf804", "https://api.mailgun.net/v3/sandboxd6d9d06b590b44d7b8dcacae8c7e1528.mailgun.org")
	// mail := mailgun.NewMessage(
	// 	"Excited User <"+cntc.Email+">",
	// 	cntc.Subject,
	// 	cntc.Msg,
	// 	"huzi8014@gmail.com",
	// 	"huzaifaa.2002@gmail.com",
	// )
	// _, _, err := mg.Send(mail)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	err := models.AddContact(cntc)
	if err != nil {
		http.Error(w, "error", http.StatusInternalServerError)
		return
	}
	flash.Msg = "Thank you for contacting me, I will get back to you shortly"

	http.Redirect(w, r, "/contact", 302)

}
