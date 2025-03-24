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
	http.Handle("/Ressources/", http.StripPrefix("/Ressources/", http.FileServer(http.Dir("./Ressources"))))

	// API
	http.HandleFunc("/login", API.Login)
	http.HandleFunc("/register", API.Register)
	http.HandleFunc("/FetchThreadPosts", API.FetchThreadPosts)
	http.HandleFunc("/UserInfo", API.UserInfo)
	http.HandleFunc("/CreateThread", API.CreateThread)
	http.HandleFunc("/CreatePost", API.CreatePost)

	// Pages
	http.HandleFunc("/", PageHandlers.LoginPage)

	// Démarage du serveur
	http.ListenAndServe(":80", nil)
}
