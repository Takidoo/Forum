package main

import (
	"Forum/API"
	"Forum/Database"
	"Forum/PageHandlers"
	"net/http"
)

func main() {
	Database.ConnectDB()
	Database.CreatePost(1, 1, "Salut")
	http.Handle("/Ressources/", http.StripPrefix("/Ressources/", http.FileServer(http.Dir("./Ressources"))))
	http.HandleFunc("/", PageHandlers.LoginPage)
	http.HandleFunc("/login", API.Login)
	http.HandleFunc("/FetchThreadPosts", API.FetchThreadPosts)
	http.HandleFunc("/GetUserInfo", API.GetUserInfo)
	http.ListenAndServe(":80", nil)
}
