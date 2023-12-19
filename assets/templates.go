package assets

import (
	"html/template"
	"log"
	"net/http"
)

type data interface{}

type TodoTemplate struct {
	tmpl   *template.Template
	layout string
}

var tmplDirs = []string{
	"templates/layouts/*.tmpl",
	"templates/partials/*.tmpl",
}

var commonTmpl = template.Must(template.ParseFS(TmplFiles, tmplDirs...))

func MakeTmpl(path string) TodoTemplate {
	cloned := template.Must(commonTmpl.Clone())
	newTmpl := template.Must(cloned.ParseFS(TmplFiles, path))

	return TodoTemplate{
		tmpl:   newTmpl,
		layout: "layout.tmpl",
	}
}

func (t TodoTemplate) Render(w http.ResponseWriter, r *http.Request, d data) {
	log.Println(t.tmpl.ExecuteTemplate(w, t.layout, d))
}
