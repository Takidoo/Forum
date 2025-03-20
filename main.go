package main

import (
	"Forum/API"
	"Forum/Database"
	"net/http"
)

func main() {
	print("test")
	Database.ConnectDB()
	http.HandleFunc("/login", API.Login)
	http.ListenAndServe(":80", nil)
}
