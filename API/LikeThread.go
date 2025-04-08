package API

import (
	"Forum/Database"
	"Forum/Forum"
	"net/http"
)

func LikeThread(w http.ResponseWriter, r *http.Request) {
	middleAuth, aerr := Database.MiddlewareAuth(w, r)
	if !middleAuth {
		http.Error(w, aerr.Error(), http.StatusUnauthorized)
		return
	}
	if r.FormValue("thread_id") == "" {
		http.Error(w, "Invalid Thread ID", http.StatusBadRequest)
		return
	}

	cookie, _ := r.Cookie("session_id")
	user, err := Forum.GetUser(cookie.Value)
	if err != nil {
		http.Error(w, "Cannot fetch user", http.StatusInternalServerError)
		return
	}

	err = Forum.AddThreadLike(r.FormValue("thread_id"), user.ID)
	if err != nil {
		http.Error(w, "Cannot like thread", http.StatusBadRequest)
		return
	}

}
