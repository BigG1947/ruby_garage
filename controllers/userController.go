package controllers

import (
	"RubyGarage_v.2.0/models"
	"RubyGarage_v.2.0/utils"
	"encoding/json"
	"log"
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request) {
	var account models.Account
	err := json.NewDecoder(r.Body).Decode(&account)
	if err != nil {
		log.Printf("%s\n", err)
		response := utils.Message(false, "Invalid request")
		w.WriteHeader(http.StatusBadRequest)
		utils.Respond(w, response)
		return
	}

	response := account.Login()
	w.WriteHeader(http.StatusOK)
	utils.Respond(w, response)
}

func Registration(w http.ResponseWriter, r *http.Request) {
	var account models.Account

	err := json.NewDecoder(r.Body).Decode(&account)
	if err != nil {
		log.Printf("%s\n", err)
		response := utils.Message(false, "Invalid request")
		w.WriteHeader(http.StatusBadRequest)
		utils.Respond(w, response)
		return
	}

	response := account.Create()
	w.WriteHeader(http.StatusOK)
	utils.Respond(w, response)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	var account models.Account

	account.Id = r.Context().Value("user").(int64)

	response := account.Get()
	w.WriteHeader(http.StatusOK)
	utils.Respond(w, response)
}
