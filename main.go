package main

import (
	"Forum/API"
	"Forum/Database"
	"Forum/PageHandlers"
	"net/http"
)

func main() {
	// Prérequis
	Database.ConnectDB()
	http.Handle("/Resources/", http.StripPrefix("/Resources/", http.FileServer(http.Dir("./Resources"))))

	// API
	http.HandleFunc("/api/login", API.Login)
	http.HandleFunc("/api/register", API.Register)
	http.HandleFunc("/api/FetchThreadPosts", API.FetchThreadPosts)
	http.HandleFunc("/api/UserInfo", API.UserInfo)
	http.HandleFunc("/api/CreateThread", API.CreateThread)
	http.HandleFunc("/api/CreatePost", API.CreatePost)
	http.HandleFunc("/api/SetUserRole", API.SetUserRole)
	http.HandleFunc("/api/DisableAccount", API.DisableAccount)

	// Pages
	http.HandleFunc("/", PageHandlers.TestPageHandler)
	http.HandleFunc("/admin", PageHandlers.AdminPageHandler)
	http.HandleFunc("/login", PageHandlers.LoginPageHandler)

	// Démarage du serveur
	http.ListenAndServe(":80", nil)
}
