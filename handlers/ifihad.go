package handlers

import (
	"fmt"
	"html/template"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.html"))
}
func Ifihad(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		tpl.ExecuteTemplate(w, "index.html", nil)
	case "/hodling":
		tpl.ExecuteTemplate(w, "hodling.html", nil)
	default:
		fmt.Fprintf(w, "You lost the way comrade!!")
	}
}
