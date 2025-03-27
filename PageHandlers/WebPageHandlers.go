package PageHandlers

import (
	"Forum/Database"
	"html/template"
	"net/http"
)

func TestPageHandler(w http.ResponseWriter, r *http.Request) {
	var tmpl, _ = template.ParseFiles("WebPages/home.html")
	tmpl.Execute(w, nil)
}
func AdminPageHandler(w http.ResponseWriter, r *http.Request) {
	if !Database.UserIsAdmin(w, r) {
		http.Error(w, "You're not admin", http.StatusUnauthorized)
		return
	}
	var tmpl, _ = template.ParseFiles("WebPages/admin.html")
	tmpl.Execute(w, nil)
}

func LoginPageHandler(w http.ResponseWriter, r *http.Request) {
	var tmpl, _ = template.ParseFiles("WebPages/login.html")
	tmpl.Execute(w, nil)
}
