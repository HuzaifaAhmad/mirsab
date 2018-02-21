package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/HuzaifaAhmad/mirsab/handlers"
	"github.com/gorilla/mux"
)

var tpl *template.Template

func main() {
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
	r := mux.NewRouter()
	r.HandleFunc("/", indexHandler)
	r.HandleFunc("/about", aboutHandler)
	r.HandleFunc("/contact", contactHandler).Methods("GET")
	r.HandleFunc("/contact", handlers.Contact).Methods("POST")
	r.HandleFunc("/login", loginHanler).Methods("GET")
	r.HandleFunc("/login", handlers.Login).Methods("POST")
	r.HandleFunc("/logout", handlers.Logout)
	r.HandleFunc("/signup", signupHanler).Methods("GET")
	r.HandleFunc("/signup", handlers.SignUp).Methods("POST")
	r.HandleFunc("/admin", handlers.Admin)
	r.HandleFunc("/admin/portfolio", handlers.Portfolio)
	r.HandleFunc("/admin/portfolio/delete", handlers.Delete).Methods("POST")
	r.HandleFunc("/admin/portfolio/upload", handlers.Upload).Methods("GET")
	r.HandleFunc("/admin/portfolio/upload", handlers.Uploader).Methods("POST")
	r.HandleFunc("/admin/portfolio/temp-post", handlers.Temp).Methods("POST")
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./static/")))
	http.Handle("/", r)
	fmt.Println("Server Started")
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

func loginHanler(w http.ResponseWriter, r *http.Request) {
	handlers.LoginHandler(w, r, tpl)
}

func signupHanler(w http.ResponseWriter, r *http.Request) {
	err := tpl.ExecuteTemplate(w, "signup.gohtml", nil)
	if err != nil {
		log.Fatalf("template execution: %s", err)
	}
}
