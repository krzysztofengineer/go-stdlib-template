package view

import (
	"embed"
	"html/template"
)

var (
	//go:embed layouts pages
	viewFS embed.FS
)

func MustNew(patterns ...string) *template.Template {
	return template.Must(template.ParseFS(viewFS, patterns...))
}
