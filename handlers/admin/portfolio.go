package admin

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"image"
	"image/jpeg"
	"io"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

// File is package level
type File struct {
	File string
}

// Imgs is package level
type Imgs struct {
	Img  string
	Rand int
}

var img Imgs

//Portfolio reads all images in static/imgs/portfolio dir and send it to the portfolio template
func Portfolio(w http.ResponseWriter, r *http.Request) {
	var imgSrc []Imgs

	root := "static/img/portfolio/"
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		img.Img = path
		imgSrc = append(imgSrc, img)
		return nil
	})

	if err != nil {
		panic(err)
	}
	for i := range imgSrc {
		imgSrc[i].Img = strings.TrimPrefix(imgSrc[i].Img, "static/")
		img.Rand = rand.Int()
	}

	err = tpl.ExecuteTemplate(w, "portfolio.gohtml", imgSrc)
	if err != nil {
		log.Fatalf("template execution: %s", err)
	}
}

//Temp holder for the Files. TODO: ADD the images to Imgs struct
func Temp(w http.ResponseWriter, r *http.Request) {
	var xf []File

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

//Uploader writes the images to files in the static/img/portfolio/ directory
func Uploader(w http.ResponseWriter, r *http.Request) {
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

//Delete deletes the image from dir
func Delete(w http.ResponseWriter, r *http.Request) {
	src := r.Body
	bs, err := ioutil.ReadAll(src)
	if err != nil {
		http.Error(w, "Error Reading File", http.StatusInternalServerError)
		return
	}

	path := string(bs)

	err = os.Remove("static/img/portfolio" + path)
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
