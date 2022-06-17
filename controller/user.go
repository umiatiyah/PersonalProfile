package controller

import (
	"main/repository"

	"github.com/gorilla/mux"
)

func UserController(r *mux.Router) {

	r.HandleFunc("/user", repository.GetUsers).Methods("GET")
	r.HandleFunc("/user/{id}", repository.GetUser).Methods("GET")
	r.HandleFunc("/user", repository.CreateUser).Methods("POST")
	r.HandleFunc("/user/{id}", repository.UpdateUser).Methods("PUT")
	r.HandleFunc("/user/{id}", repository.DeleteUser).Methods("DELETE")

}
