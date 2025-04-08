package API

import (
	"Forum/Database"
	"net/http"
)

func LikeThread(w http.ResponseWriter, r *http.Request) {
	middleAuth, aerr := Database.MiddlewareAuth(w, r)
	if !middleAuth {
		http.Error(w, aerr.Error(), http.StatusUnauthorized)
		return
	}
}
