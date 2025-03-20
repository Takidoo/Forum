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
	http.HandleFunc("/", PageHandlers.LoginPage)
	http.HandleFunc("/login", API.Login)
	http.ListenAndServe(":80", nil)
}
