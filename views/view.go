package views

import (
	"html/template"
)

type View struct {
	Template *template.Template
	Layout string
}

// NewView creates a new view with common layout files
// attached as well to the resulting View
func NewView(layout string, files ...string) *View {
	// all new files will add the footer
	files = append(files, 
		"views/layouts/footer.gohtml",
		"views/layouts/bootstrap.gohtml",
		"views/layouts/navbar.gohtml")
	t, err := template.ParseFiles(files...)
	if err != nil {
		panic(err)
	}

	return &View {
		Template: t,
		Layout: layout,
	}
}