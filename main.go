package main

import (
	"crud-app/db"
	"crud-app/handlers"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	db.Init()

	r := mux.NewRouter()

	r.HandleFunc("/users", handlers.GetUsers).Methods("GET")
	r.HandleFunc("/users", handlers.CreateUser).Methods("POST")
	r.HandleFunc("/users/{id}", handlers.GetUser).Methods("GET")
	r.HandleFunc("/users/{id}", handlers.UpdateUser).Methods("PUT")
	r.HandleFunc("/users/{id}", handlers.DeleteUser).Methods("DELETE")

	log.Println("Server running on :8000")
	log.Fatal(http.ListenAndServe(":8000", r))
}
