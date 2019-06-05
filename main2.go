package main

import (
	"github.com/gorilla/mux"
	"html/template"
	"net/http"
)

var homeTemplate *template.Template

func main() {
	var err error
	homeTemplate, err = template.ParseFiles("views/index.gohtml")
	if err != nil {
		panic(err)
	}
	r := mux.NewRouter()
	r.HandleFunc("/", index)

	// Styles
	assetHandler := http.FileServer(http.Dir("./public/dist/"))
	assetHandler = http.StripPrefix("/public/dist/", assetHandler)
	r.PathPrefix("/public/dist/").Handler(assetHandler)

	// JS
	jsHandler := http.FileServer(http.Dir("./public/dist/"))
	jsHandler = http.StripPrefix("/public/dist/", jsHandler)
	r.PathPrefix("/public/dist/").Handler(jsHandler)

	//Images
	imageHandler := http.FileServer(http.Dir("./public/dist/images/"))
	r.PathPrefix("/public/dist/images/").Handler(http.StripPrefix("/public/dist/images/", imageHandler))

	http.ListenAndServe(":3050", r)
}

func index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	if err := homeTemplate.Execute(w, nil); err != nil {
		panic(err)
	}
}
