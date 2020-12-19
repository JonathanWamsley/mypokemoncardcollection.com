package main

import (
	"net/http"

	"mypokemoncardcollection.com/views"

	"github.com/gorilla/mux"
)

var homeView *views.View
var aboutView *views.View

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	err := homeView.Template.ExecuteTemplate(w,
		homeView.Layout, nil)
	if err != nil {
	  panic(err)
	}
  }
  
  func about(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	err := aboutView.Template.ExecuteTemplate(w,
		aboutView.Layout, nil)
	if err != nil {
	  panic(err)
	}
  }

func main() {
	homeView = views.NewView("bootstrap", "views/home.gohtml")
	aboutView = views.NewView("bootstrap" ,"views/about.gohtml")

	r := mux.NewRouter()
	r.HandleFunc("/", home)
	r.HandleFunc("/about", about)
	http.ListenAndServe(":3000", r)
}