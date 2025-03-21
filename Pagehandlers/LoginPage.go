package Pagehandlers

import (
	"html/template"
	"net/http"
)

func LoginPage(w http.ResponseWriter, r *http.Request) {
	var tmpl, _ = template.ParseFiles("WebPages/connexion.html")
	tmpl.Execute(w, nil)
}
