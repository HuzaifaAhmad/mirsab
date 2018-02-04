package main

import (
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var tpl *template.Template

func main() {
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
	r := mux.NewRouter()
	r.HandleFunc("/", indexHandler)
	r.HandleFunc("/about", aboutHandler)
	r.HandleFunc("/contact", contactHandler)
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./static/")))
	http.Handle("/", r)
	http.ListenAndServe(":8080", nil)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	err := tpl.ExecuteTemplate(w, "home.gohtml", nil)
	if err != nil {
		log.Fatalf("template execution: %s", err)
	}
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	err := tpl.ExecuteTemplate(w, "about.gohtml", nil)
	if err != nil {
		log.Fatalf("template execution: %s", err)
	}
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	err := tpl.ExecuteTemplate(w, "contact.gohtml", nil)
	if err != nil {
		log.Fatalf("template execution: %s", err)
	}
}
