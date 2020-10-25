package controllers

import (
	"html/template"
	"log"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	temp, err := template.ParseFiles("templates/index.html")
	if err != nil {
		log.Printf("%s\n", err)
		w.WriteHeader(500)
		return
	}
	if err := temp.Execute(w, nil); err != nil {
		log.Printf("%s\n", err)
		w.WriteHeader(500)
		return
	}
}
