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
	http.HandleFunc("/register", API.Login)
	http.HandleFunc("/FetchThreadPosts", API.FetchThreadPosts)
	http.HandleFunc("/GetUserInfo", API.GetUserInfo)
	http.HandleFunc("/CreateThread", API.CreateThread)

	// Pages
	http.HandleFunc("/", PageHandlers.LoginPage)

	// Démarage du serveur
	http.ListenAndServe(":80", nil)
}
