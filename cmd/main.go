package main

import (
	"log"
	"net/http"

	"rest-api/config"
	"rest-api/controllers"

	"github.com/gorilla/mux"
)

func main() {
	config.InitDB()

	r := mux.NewRouter()

	r.HandleFunc("/signup", controllers.SignupHandler).Methods("POST")
	r.HandleFunc("/login", controllers.LoginHandler).Methods("POST")

	r.HandleFunc("/users", controllers.GetUsersHandler).Methods("GET")
	r.HandleFunc("/users/{username}", controllers.DeleteUserHandler).Methods("DELETE")

	log.Println("Server started at :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
