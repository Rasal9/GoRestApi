package main

import (
	"log"
	"net/http"

	"rest-api/config"
	"rest-api/controllers"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {
	config.InitDB()

	r := mux.NewRouter()

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"},
		AllowedMethods:   []string{"GET", "POST", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
		MaxAge:           300,
	})

	handler := c.Handler(r)

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Rest-api"))
	}).Methods("GET")
	r.HandleFunc("/signup", controllers.SignupHandler).Methods("POST")
	r.HandleFunc("/login", controllers.LoginHandler).Methods("POST")
	r.HandleFunc("/users", controllers.GetUsersHandler).Methods("GET")
	r.HandleFunc("/users/{username}", controllers.DeleteUserHandler).Methods("DELETE")

	log.Println("Server started at :8080")
	log.Fatal(http.ListenAndServe(":8080", handler))
}
