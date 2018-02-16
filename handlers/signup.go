package handlers

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/HuzaifaAhmad/mirsab/models"
	"golang.org/x/crypto/bcrypt"
)

func SignUp(w http.ResponseWriter, r *http.Request) {
	fname := r.FormValue("fname")
	lname := r.FormValue("lname")
	email := r.FormValue("email")
	pass := r.FormValue("pass")
	repass := r.FormValue("repass")

	fname, lname, email = toCase(fname, lname, email)

	if pass != repass {
		fmt.Fprintln(w, "Pass donsnt Match")
	}

	enPass, err := encryptPassword(pass)
	if err != nil {
		fmt.Fprintln(w, "Internal Server Error")
	}

	err = models.AddUser(fname, lname, email, enPass)
	if err != nil {
		fmt.Fprintln(w, "Error: Could n user")
	} else {
		fmt.Fprintln(w, "Successfuly added user")
	}

}

func encryptPassword(pass string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	if err != nil {
		// TODO: Properly handle error
		return "", err
	}

	toStr := string(hash)

	return toStr, nil
}

func toCase(fn, ln, eml string) (fname, lname, email string) {
	fname = strings.Title(strings.ToLower(fn))
	lname = strings.Title(strings.ToLower(ln))
	email = strings.ToLower(eml)

	return fname, lname, email
}
