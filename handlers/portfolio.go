package handlers

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"image"
	"image/jpeg"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

type File struct {
	File string
}

type Imgs struct {
	Img []string
}

func Portfolio(w http.ResponseWriter, r *http.Request) {
	email := getSession(r)
	if email == "" {
		http.Redirect(w, r, "/login", 302)
		return
	}
	var imgSrc []string

	root := "static/img/portfolio/"
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		imgSrc = append(imgSrc, path)
		return nil
	})

	if err != nil {
		panic(err)
	}
	for i := range imgSrc {
		imgSrc[i] = strings.TrimPrefix(imgSrc[i], "static/")
	}

	err = tpl.ExecuteTemplate(w, "portfolio.gohtml", Imgs{
		Img: imgSrc,
	})
	if err != nil {
		log.Fatalf("template execution: %s", err)
	}
}

func Temp(w http.ResponseWriter, r *http.Request) {

	var xf []File
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

	fn := File{
		File: en,
	}

	xf = append(xf, fn)
	_, err = json.Marshal(xf)
	if err != nil {
		fmt.Println(err)
	}
	io.WriteString(w, en)
}

func Uploader(w http.ResponseWriter, r *http.Request) {
	email := getSession(r)
	if email == "" {
		http.Redirect(w, r, "/login", http.StatusNonAuthoritativeInfo)
		return
	}

	src := r.Body
	if src == nil {
		fmt.Println("Empty")
		return
	}

	bs, err := ioutil.ReadAll(src)
	if err != nil {
		http.Error(w, "Error Reading File", http.StatusInternalServerError)
		return
	}

	var xf []string

	err = json.Unmarshal(bs, &xf)
	if err != nil {
		fmt.Println(err)
	}

	for i := 0; i < len(xf); i++ {
		data := xf[i]
		reader := base64.NewDecoder(base64.StdEncoding, strings.NewReader(data))
		m, _, err := image.Decode(reader)
		if err != nil {
			log.Fatal(err)
		}

		//Encode from image format to writer
		num, err := getPicNum()
		if err != nil {
			http.Error(w, "Error Writing image file.", http.StatusInternalServerError)
			return
		}
		pngFilename := "static/img/portfolio/pic" + strconv.Itoa(num) + ".jpg"
		f, err := os.OpenFile(pngFilename, os.O_WRONLY|os.O_CREATE, 0777)
		if err != nil {
			log.Fatal(err)
			return
		}

		err = jpeg.Encode(f, m, &jpeg.Options{Quality: 75})
		if err != nil {
			log.Fatal(err)
			return
		}
		// num++
	}
}

func Delete(w http.ResponseWriter, r *http.Request) {
	email := getSession(r)
	if email == "" {
		http.Redirect(w, r, "/login", http.StatusNonAuthoritativeInfo)
		return
	}

	src := r.Body
	bs, err := ioutil.ReadAll(src)
	if err != nil {
		http.Error(w, "Error Reading File", http.StatusInternalServerError)
		return
	}

	toString := string(bs)
	prefix := "http://localhost:8080"
	ns := strings.TrimPrefix(toString, prefix)

	err = os.Remove("static" + ns)
	if err != nil {
		http.Error(w, "Error Deleting File", http.StatusInternalServerError)
		log.Println(err)
	}
}

func getPicNum() (int, error) {
	var files []string

	root := "static/img/portfolio"
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		files = append(files, path)
		return nil
	})
	if err != nil {
		panic(err)
	}
	var ia int
	for _, file := range files {
		p := strings.TrimPrefix(file, "static/img/portfolio/pic")
		s := strings.TrimSuffix(p, ".jpg")
		if s == "" {
			return 0, nil
		}
		i, err := strconv.Atoi(s)
		if err != nil {
			return 0, err
		}
		ia = i + 1

	}
	return ia, nil
}
