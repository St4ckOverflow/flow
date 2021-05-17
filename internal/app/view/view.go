package view

import (
	"embed"
	"log"
	"net/http"
	"path"
	"text/template"
)

//go:embed html/index.html
var htmlFiles embed.FS

func htmlPath(name string) string {
	return path.Join("html", name)
}

func Show(rw http.ResponseWriter, name string, data interface{}) {
	tpl, err := template.ParseFS(htmlFiles, htmlPath(name))
	if err != nil {
		log.Println(err)
	}
	tpl.Execute(rw, data)
}
