package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func main() {
	// handle CSS file requests
	css := http.FileServer(http.Dir("./css"))
	http.Handle("/css/", http.StripPrefix("/css", css))

	// handle Image file requests
	imgs := http.FileServer(http.Dir("./images"))
	http.Handle("/images/", http.StripPrefix("/images", imgs))

	// handle Scripts file requests
	scripts := http.FileServer(http.Dir("./scripts"))
	http.Handle("/scripts/", http.StripPrefix("/scripts", scripts))

	// page renders
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		render(w, "home.page.gohtml")
	})

	http.HandleFunc("/computer", func(w http.ResponseWriter, r *http.Request) {
		render(w, "computer.page.gohtml")
	})

	log.Println("Starting frontend service on port 80")
	if err := http.ListenAndServe(":80", nil); err != nil {
		log.Panic(err)
	}
}

func render(w http.ResponseWriter, t string) {
	partials := []string{
		"./templates/base.layout.gohtml",
		"./templates/header.partial.gohtml",
		"./templates/sidebar.partial.gohtml",
	}

	var templateSlice []string
	templateSlice = append(templateSlice, fmt.Sprintf("./templates/%s", t))

	for _, x := range partials {
		templateSlice = append(templateSlice, x)
	}

	tmpl, err := template.ParseFiles(templateSlice...)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := tmpl.Execute(w, nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
