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
			"bootstrap", "static/home"),
		About: views.NewView(
			"bootstrap", "static/about"),
	}
}
