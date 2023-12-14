package assets

import (
	"html/template"
)

var tmplDirs = []string{
	"templates/layouts/*.tmpl",
	"templates/partials/*.tmpl",
}

var commonTmpl = template.Must(template.ParseFS(TmplFiles, tmplDirs...))

func MakeTmpl(path string) *template.Template {
	cloned := template.Must(commonTmpl.Clone())
	return template.Must(cloned.ParseFS(TmplFiles, path))
}
