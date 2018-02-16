package handlers

import (
	"encoding/base64"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

func Portfolio(w http.ResponseWriter, r *http.Request) {
	email := getSession(r)
	if email != "" {
		err := tpl.ExecuteTemplate(w, "portfolio.gohtml", nil)
		if err != nil {
			log.Fatalf("template execution: %s", err)
		}
	} else {
		http.Redirect(w, r, "/login", 302)
	}
}

func Temp(w http.ResponseWriter, r *http.Request) {
	email := getSession(r)
	if email == "" {
		http.Redirect(w, r, "/login", http.StatusNonAuthoritativeInfo)
		return
	}

	src, _, err := r.FormFile("file")
	if src == nil {
		return
	}
	if err != nil {
		http.Error(w, "Error Uploading File", http.StatusInternalServerError)
		fmt.Println(err)
		return
	}

	defer src.Close()

	bs, err := ioutil.ReadAll(src)
	if err != nil {
		http.Error(w, "Error Reading File", http.StatusInternalServerError)
		return
	}

	en := base64.StdEncoding.EncodeToString(bs)

	io.WriteString(w, en)
}
