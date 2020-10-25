package main

import (
	"RubyGarage_v.2.0/app"
	"RubyGarage_v.2.0/controllers"
	"RubyGarage_v.2.0/models"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"log"
	"net/http"
)

func main() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	models.InitDB()

	router := mux.NewRouter()
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static", http.FileServer(http.Dir("static"))))
	router.HandleFunc("/", controllers.Index).Methods(http.MethodGet)

	router.HandleFunc("/user/login", controllers.Login).Methods(http.MethodPost)
	router.HandleFunc("/user/registration", controllers.Registration).Methods(http.MethodPost)
	router.HandleFunc("/user", controllers.GetUser).Methods(http.MethodGet)

	router.HandleFunc("/project/create", controllers.CreateProject).Methods(http.MethodPost)
	router.HandleFunc("/projects", controllers.GetUserProjects).Methods(http.MethodGet)
	router.HandleFunc("/project/delete", controllers.DeleteProject).Methods(http.MethodDelete)
	router.HandleFunc("/project/edit", controllers.EditProject).Methods(http.MethodPut)

	router.HandleFunc("/task/create", controllers.CreateTask).Methods(http.MethodPost)
	router.HandleFunc("/task/edit", controllers.EditTask).Methods(http.MethodPut)
	router.HandleFunc("/task/delete", controllers.DeleteTask).Methods(http.MethodDelete)
	router.HandleFunc("/task/check", controllers.CheckTask).Methods(http.MethodPut)
	router.HandleFunc("/task/priority/up", controllers.UpPriorityTask).Methods(http.MethodPost)
	router.HandleFunc("/task/priority/down", controllers.DownPriorityTask).Methods(http.MethodPost)

	router.Use(app.JwtAuthentication)

	log.Fatal(http.ListenAndServe(":9000", router))
}
