package views

import (
	"html/template"
)

type View struct {
	Template *template.Template
}

// NewView creates a new view with common layout files
// attached as well to the resulting View
func NewView(files ...string) *View {
	// all new files will add the footer
	files = append(files, "views/layouts/footer.gohtml")
	t, err := template.ParseFiles(files...)
	if err != nil {
		panic(err)
	}

	return &View {
		Template: t,
	}
}