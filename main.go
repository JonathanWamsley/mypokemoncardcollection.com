package main

import (
	"net/http"

	"mypokemoncardcollection.com/controllers"

	"github.com/gorilla/mux"
)

func main() {
	staticController := controllers.NewStatic()
	usersController := controllers.NewUsers()
	
	r := mux.NewRouter()
	r.Handle("/", staticController.Home).Methods("GET")
	r.Handle("/about", staticController.About).Methods("GET")
	r.HandleFunc("/signup", usersController.New).Methods("GET")
	r.HandleFunc("/signup", usersController.Create).Methods("POST")
	http.ListenAndServe(":3000", r)
}