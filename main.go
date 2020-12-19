package main

import (
	"net/http"

	"mypokemoncardcollection.com/views"

	"github.com/gorilla/mux"
)

var homeView *views.View
var aboutView *views.View

// A helper function that panics of any error
func must(err error) {
	if err != nil {
		panic(err)
	}
}

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	must(homeView.Render(w, nil))
  }
  
  func about(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	must(aboutView.Render(w, nil))
  }

func main() {
	homeView = views.NewView("bootstrap", "views/home.gohtml")
	aboutView = views.NewView("bootstrap" ,"views/about.gohtml")

	r := mux.NewRouter()
	r.HandleFunc("/", home)
	r.HandleFunc("/about", about)
	http.ListenAndServe(":3000", r)
}