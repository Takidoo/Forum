package main

import (
	"Forum/API"
	"Forum/Database"
	"Forum/PageHandlers"
	"net/http"
)

func main() {
	Database.ConnectDB()
	var err = Database.RegisterUser("Takido", "test")
	print(err)
	http.HandleFunc("/", PageHandlers.LoginPage)
	http.HandleFunc("/login", API.Login)
	http.ListenAndServe(":80", nil)
}
