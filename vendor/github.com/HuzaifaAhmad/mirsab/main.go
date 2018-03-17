package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/HuzaifaAhmad/mirsab/handlers"
	"github.com/HuzaifaAhmad/mirsab/handlers/admin"
	"github.com/codegangsta/negroni"
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
	r.HandleFunc("/login", admin.Login).Methods("POST")
	r.HandleFunc("/logout", admin.Logout)
	r.HandleFunc("/signup", signupHanler).Methods("GET")
	r.HandleFunc("/signup", handlers.SignUp).Methods("POST")

	adm := mux.NewRouter()

	r.PathPrefix("/admin").Handler(negroni.New(
		negroni.HandlerFunc(auth),
		negroni.Wrap(adm),
	))

	ad := adm.PathPrefix("/admin").Subrouter()
	ad.Path("").HandlerFunc(admin.Dashboard).Methods("GET")
	ad.Path("/portfolio").HandlerFunc(admin.Portfolio).Methods("GET")
	ad.Path("/portfolio/delete").HandlerFunc(admin.Delete).Methods("POST")
	ad.Path("/portfolio/upload").HandlerFunc(admin.Upload).Methods("GET")
	ad.Path("/portfolio/upload").HandlerFunc(admin.Uploader).Methods("POST")
	ad.Path("/portfolio/temp-post").HandlerFunc(admin.Temp).Methods("POST")
	ad.Path("/contact").HandlerFunc(admin.Contact).Methods("GET")
	ad.Path("/contact/{id:[0-9]+}").HandlerFunc(admin.ContactDetails).Methods("GET")

	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./static/")))

	// http.Handle("/", r)
	fmt.Println("Server Started")
	http.ListenAndServe(":8080", r)
}

func auth(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	email := admin.GetSession(r)
	if email == "" {
		http.Redirect(w, r, "/login", 302)
		return
	}

	next(w, r)

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
	handlers.ContactPage(w, r, tpl)
}

func loginHanler(w http.ResponseWriter, r *http.Request) {
	admin.LoginHandler(w, r, tpl)
}

func signupHanler(w http.ResponseWriter, r *http.Request) {
	err := tpl.ExecuteTemplate(w, "signup.gohtml", nil)
	if err != nil {
		log.Fatalf("template execution: %s", err)
	}
}
