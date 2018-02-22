package admin

import (
	"database/sql"
	"html/template"
	"log"
	"net/http"

	"github.com/HuzaifaAhmad/mirsab/models"
	"github.com/gorilla/securecookie"
	"golang.org/x/crypto/bcrypt"
)

// User is package level
type User struct {
	ID        int
	FirstName string
	LastName  string
	Email     string
	Password  string
}

type page struct {
	Title string
	Err   string
}

var p page
var cookieHandler = securecookie.New(securecookie.GenerateRandomKey(64), securecookie.GenerateRandomKey(32))

//LoginHandler Renders the login page template
func LoginHandler(w http.ResponseWriter, r *http.Request, tpl *template.Template) {
	err := tpl.ExecuteTemplate(w, "login.gohtml", p)
	if err != nil {
		log.Fatalf("template execution: %s", err)
	}
}

// Login takes post req and checks the credentilas against the db and stores it in User struct
func Login(w http.ResponseWriter, r *http.Request) {
	email := r.FormValue("email")
	pass := r.FormValue("pass")
	redirectTarget := "/login"
	usr, err := models.GetUser(email, pass)

	if err == sql.ErrNoRows {
		p.Err = "Incorrect email or password"

	} else if err != nil {
		str := "Internal Server Error"
		p.Err = str
		log.Println(str)
	}

	u := User{
		ID:        usr.ID,
		FirstName: usr.FirstName,
		LastName:  usr.LastName,
		Email:     usr.Email,
		Password:  usr.Password,
	}

	ok := checkPass(pass, u.Password)
	if ok {
		redirectTarget = "/admin"
		setSession(u.Email, w)
	}

	http.Redirect(w, r, redirectTarget, 302)
}

// Logout deletes the session
func Logout(w http.ResponseWriter, r *http.Request) {
	clearSession(w)
	http.Redirect(w, r, "/login", 302)
}

func checkPass(userPass, dbPass string) bool {
	hashFromDatabase := dbPass

	// Comparing the password with the hash
	if err := bcrypt.CompareHashAndPassword([]byte(hashFromDatabase), []byte(userPass)); err != nil {
		return false
	}

	return true
}

func setSession(email string, w http.ResponseWriter) {
	value := map[string]string{
		"name": email,
	}
	if encoded, err := cookieHandler.Encode("session", value); err == nil {
		cookie := &http.Cookie{
			Name:  "session",
			Value: encoded,
			Path:  "/",
		}
		http.SetCookie(w, cookie)
	}
}

func clearSession(w http.ResponseWriter) {
	cookie := &http.Cookie{
		Name:   "session",
		Value:  "",
		Path:   "/",
		MaxAge: -1,
	}

	http.SetCookie(w, cookie)
}
