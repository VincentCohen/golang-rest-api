package main

import (
	"net/http"
	
	"github.com/gorilla/mux"
)


func main () {
	r := mux.NewRouter()
	//
	//ur := repository.NewUserRepository(getDatabase())
	//uc := controllers.NewUserController(ur)
	//
	//r.GET("/users", uc.GetUsers)

	r.HandleFunc("/api/dogs", ReadDogs).Methods("GET")
	r.HandleFunc("/api/dogs", CreateDog).Methods("POST")
	r.HandleFunc("/api/dogs", UpdateDog).Methods("PUT")
	r.HandleFunc("/api/dogs", DeleteDog).Methods("DELETE")

	http.ListenAndServe(":3000", r)
}

func ReadDogs() {}
func CreateDog() {}
func UpdateDog() {}
func DeleteDog() {}
