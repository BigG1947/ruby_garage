package controllers

import (
	"RubyGarage_v.2.0/models"
	"RubyGarage_v.2.0/utils"
	"encoding/json"
	"log"
	"net/http"
)

func CreateProject(w http.ResponseWriter, r *http.Request) {
	var p models.Project

	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		log.Printf("%s\n", err)
		response := utils.Message(false, "Invalid request data")
		w.WriteHeader(http.StatusBadRequest)
		utils.Respond(w, response)
		return
	}

	p.IdUser = r.Context().Value("user").(int64)

	response := p.Create()
	w.WriteHeader(http.StatusOK)
	utils.Respond(w, response)
}

func EditProject(w http.ResponseWriter, r *http.Request) {
	var p models.Project

	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		log.Printf("%s\n", err)
		response := utils.Message(false, "Invalid request data")
		w.WriteHeader(http.StatusBadRequest)
		utils.Respond(w, response)
		return
	}

	p.IdUser = r.Context().Value("user").(int64)

	response := p.Edit()
	w.WriteHeader(http.StatusOK)
	utils.Respond(w, response)
}

func DeleteProject(w http.ResponseWriter, r *http.Request) {
	var p models.Project

	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		log.Printf("%s\n", err)
		response := utils.Message(false, "Invalid request data")
		w.WriteHeader(http.StatusBadRequest)
		utils.Respond(w, response)
		return
	}

	p.IdUser = r.Context().Value("user").(int64)

	response := p.Delete()

	w.WriteHeader(http.StatusOK)
	utils.Respond(w, response)
}

func GetUserProjects(w http.ResponseWriter, r *http.Request) {
	pl := models.ProjectList{}

	uid := r.Context().Value("user").(int64)

	response := pl.Get(uid)

	w.WriteHeader(http.StatusOK)
	utils.Respond(w, response)
}
