package PageHandlers

import (
	"Forum/Forum"
	"html/template"
	"net/http"
)

type HomePageData struct {
	Username      string
	LastedThreads []Forum.Thread
	IsLogged      bool
	IsAdmin       bool
}

type AdminPageData struct {
	Username string
}

func TestPageHandler(w http.ResponseWriter, r *http.Request) {
	var tmpl, _ = template.ParseFiles("WebPages/test.html")
	tmpl.Execute(w, nil)
}
func HomePageHandler(w http.ResponseWriter, r *http.Request) {
	var tmpl, _ = template.ParseFiles("WebPages/home.html")
	cookie, err := r.Cookie("session_id")
	if err == nil {
		user, err := Forum.GetUser(cookie.Value)
		if err == nil {
			tmpl.Execute(w, HomePageData{
				Username:      user.Username,
				LastedThreads: Forum.GetLastedThreads(10),
				IsLogged:      true,
				IsAdmin:       user.Role == 2,
			})
			return
		}
	}

	tmpl.Execute(w, HomePageData{
		Username:      "",
		LastedThreads: Forum.GetLastedThreads(10),
		IsLogged:      false,
		IsAdmin:       false,
	})
}
func AdminPageHandler(w http.ResponseWriter, r *http.Request) {
	if !Forum.UserIsAdmin(w, r) {
		http.Error(w, "You're not admin", http.StatusUnauthorized)
		return
	}
	cookie, _ := r.Cookie("session_id")
	user, _ := Forum.GetUser(cookie.Value)
	var tmpl, _ = template.ParseFiles("WebPages/admin.html")
	tmpl.Execute(w, AdminPageData{
		Username: user.Username,
	})
}

func LoginPageHandler(w http.ResponseWriter, r *http.Request) {
	var tmpl, _ = template.ParseFiles("WebPages/login.html")
	tmpl.Execute(w, nil)
}
