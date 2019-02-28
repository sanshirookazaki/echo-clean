package template

import (
	"html/template"
	"io"
)

type Template struct {
	templates *template.Template
}

func NewTemplate(path string) *Template {
	return &Template{
		templates: template.Must(template.ParseGlob("views/*.html")),
	}
}

func (t *Template) Render(w io.Writer, name string, data interface{}) error {
	return t.templates.ExecuteTemplate(w, name, data)
}
