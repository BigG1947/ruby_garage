package controllers

import (
	"RubyGarage_v.2.0/models"
	"RubyGarage_v.2.0/utils"
	"encoding/json"
	"log"
	"net/http"
)

func CreateTask(w http.ResponseWriter, r *http.Request) {
	var t models.Task

	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		log.Printf("%s\n", err)
		w.WriteHeader(http.StatusBadRequest)
		response := utils.Message(false, "Invalid request data")
		utils.Respond(w, response)
		return
	}
	uid := r.Context().Value("user").(int64)
	response := t.Create(uid)

	w.WriteHeader(200)
	utils.Respond(w, response)
}

func EditTask(w http.ResponseWriter, r *http.Request) {
	var t models.Task

	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		log.Printf("%s\n", err)
		w.WriteHeader(http.StatusBadRequest)
		response := utils.Message(false, "Invalid request data")
		utils.Respond(w, response)
		return
	}

	uid := r.Context().Value("user").(int64)

	response := t.Edit(uid)
	w.WriteHeader(http.StatusOK)
	utils.Respond(w, response)
}

func DeleteTask(w http.ResponseWriter, r *http.Request) {
	var t models.Task

	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		log.Printf("%s\n", err)
		w.WriteHeader(http.StatusBadRequest)
		response := utils.Message(false, "Invalid request data")
		utils.Respond(w, response)
		return
	}

	uid := r.Context().Value("user").(int64)

	response := t.Delete(uid)
	w.WriteHeader(http.StatusOK)
	utils.Respond(w, response)
}

func CheckTask(w http.ResponseWriter, r *http.Request) {
	var t models.Task

	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		log.Printf("%s\n", err)
		w.WriteHeader(http.StatusBadRequest)
		response := utils.Message(false, "Invalid request data")
		utils.Respond(w, response)
		return
	}

	uid := r.Context().Value("user").(int64)

	response := t.Check(uid)
	w.WriteHeader(http.StatusOK)
	utils.Respond(w, response)
}

func UpPriorityTask(w http.ResponseWriter, r *http.Request) {
	var t models.Task

	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		log.Printf("%s\n", err)
		w.WriteHeader(http.StatusBadRequest)
		response := utils.Message(false, "Invalid request data")
		utils.Respond(w, response)
		return
	}

	uid := r.Context().Value("user").(int64)

	response := t.PriorityUp(uid)
	w.WriteHeader(http.StatusOK)
	utils.Respond(w, response)
}

func DownPriorityTask(w http.ResponseWriter, r *http.Request) {
	var t models.Task

	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		log.Printf("%s\n", err)
		w.WriteHeader(http.StatusBadRequest)
		response := utils.Message(false, "Invalid request data")
		utils.Respond(w, response)
		return
	}

	uid := r.Context().Value("user").(int64)

	response := t.PriorityDown(uid)
	w.WriteHeader(http.StatusOK)
	utils.Respond(w, response)
}
