package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/", nil)
	http.ListenAndServe(":80", nil)
}
