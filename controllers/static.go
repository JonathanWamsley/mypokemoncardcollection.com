package controllers

import (
	"mypokemoncardcollection.com/views"
)

type Static struct {
	Home *views.View
	About *views.View
}

func NewStatic() *Static {
	return &Static{
		Home: views.NewView(
			"bootstrap", "views/static/home.gohtml"),
		About: views.NewView(
			"bootstrap", "views/static/about.gohtml"),
	}
}
