package main

import (
	"Forum/API"
	"Forum/Database"
	"Forum/PageHandlers"
	"net/http"
)

func main() {
	Database.ConnectDB()
	Database.RegisterUser("Takido", "test")
	http.Handle("/Ressources/", http.StripPrefix("/Ressources/", http.FileServer(http.Dir("./Ressources"))))
	http.HandleFunc("/", PageHandlers.LoginPage)
	http.HandleFunc("/login", API.Login)
	http.ListenAndServe(":80", nil)
}
