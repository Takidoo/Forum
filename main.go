package main

import (
	"Forum/API"
	"Forum/Database"
	"Forum/PageHandlers"
	"net/http"
)

func main() {
	Database.ConnectDB()
<<<<<<< HEAD
	var err = Database.RegisterUser("Takido", "test")
	print(err)
=======
	Database.RegisterUser("Takido", "test")
>>>>>>> 3556f0aece8406937bfaf1dbff4d4b0c0a91241f
	http.HandleFunc("/", PageHandlers.LoginPage)
	http.HandleFunc("/login", API.Login)
	http.ListenAndServe(":80", nil)
}
