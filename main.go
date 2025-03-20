package main

import (
	"Forum/API"
	"net/http"
)

func main() {
	print("test")
	http.HandleFunc("/login", API.Login)
	http.ListenAndServe(":80", nil)
}
