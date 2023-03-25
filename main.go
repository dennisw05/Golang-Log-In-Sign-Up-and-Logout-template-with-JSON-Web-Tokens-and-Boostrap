package main

import (
	"log"
	"loginRegisterTemplate/handlers"
	"loginRegisterTemplate/helpers"
	"loginRegisterTemplate/repositories"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	if os.Getenv("APP_MODE") != "production" {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
		}
	}

	if repositories.Connected() {
		log.Println("Started database connection")
	}

	r := mux.NewRouter()

	r.Use(helpers.CheckAuth)

	r.HandleFunc("/login", handlers.LoginHandler).Methods("GET")
	r.HandleFunc("/login", handlers.LoginPostHandler).Methods("POST")
	r.HandleFunc("/signup", handlers.SignupHandler).Methods("GET")
	r.HandleFunc("/signup", handlers.SignupPostHandler).Methods("POST")
	r.HandleFunc("/logout", handlers.LogoutHandler).Methods("GET")
	r.HandleFunc("/", handlers.HomeHandler).Methods("GET")

	r.PathPrefix("/public/").Handler(http.StripPrefix("/public/", http.FileServer(http.Dir("public"))))

	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), r))
}
