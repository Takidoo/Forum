package main

import (
	"Forum/API"
	"Forum/Database"
	"Forum/Pagehandlers"
	"net/http"
)

func main() {
	Database.ConnectDB()
	Database.RegisterUser("Takido", "test")
	http.HandleFunc("/", Pagehandlers.LoginPage)
	http.HandleFunc("/login", API.Login)
	http.ListenAndServe(":80", nil)
}
