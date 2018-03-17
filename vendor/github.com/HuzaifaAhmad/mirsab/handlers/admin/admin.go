package admin

import (
	"html/template"
	"log"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/admin/*.gohtml"))
}

//Dashboard renders dashboard template
func Dashboard(w http.ResponseWriter, r *http.Request) {
	err := tpl.ExecuteTemplate(w, "dashboard.gohtml", nil)
	if err != nil {
		log.Fatalf("template execution: %s", err)
	}

}

// Upload renders upload template
func Upload(w http.ResponseWriter, r *http.Request) {
	err := tpl.ExecuteTemplate(w, "upload.gohtml", nil)
	if err != nil {
		log.Fatalf("template execution: %s", err)
	}
}

func GetSession(r *http.Request) (email string) {
	if cookie, err := r.Cookie("session"); err == nil {
		cookieValue := make(map[string]string)
		if err = cookieHandler.Decode("session", cookie.Value, &cookieValue); err == nil {
			email = cookieValue["name"]
		}
	}
	return email
}
