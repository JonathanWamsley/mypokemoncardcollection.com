package main

import (
	"net/http"
	"fmt"

	"mypokemoncardcollection.com/controllers"
	"mypokemoncardcollection.com/models"

	"github.com/gorilla/mux"
)

// all this is dummy config
const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "password"
	dbname   = "mypokemoncardcollection_dev"
  )

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	us, err := models.NewUserService(psqlInfo)
	if err != nil {
		panic(err)
	}
	defer us.Close()
	us.AutoMigrate()
	
	staticController := controllers.NewStatic()
	usersController := controllers.NewUsers()
	
	r := mux.NewRouter()
	r.Handle("/", staticController.Home).Methods("GET")
	r.Handle("/about", staticController.About).Methods("GET")
	r.HandleFunc("/signup", usersController.New).Methods("GET")
	r.HandleFunc("/signup", usersController.Create).Methods("POST")
	http.ListenAndServe(":3000", r)
}