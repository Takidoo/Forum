package API

import (
	"Forum/Database"
	"encoding/json"
	"net/http"
)

func CreatePost(w http.ResponseWriter, r *http.Request) {
	if !Database.MiddlewareAuth(w, r) {
		return
	}

	thread_id := r.FormValue("thread_id")
	message := r.FormValue("message")
	if message == "" || thread_id == "" {
		http.Error(w, "Invalid Args", http.StatusBadRequest)
	}

	cookie, _ := r.Cookie("session_id")
	var user User
	resp, _ := http.Get("http://127.0.0.1/GetUserInfo?session=" + cookie.Value)
	json.NewDecoder(resp.Body).Decode(&user)
	_, qerr := Database.DB.Exec(`INSERT INTO posts (thread_id, user_id, content) VALUES (?, ?, ?)`, thread_id, user.ID, message)
	if qerr != nil {
		http.Error(w, "Cannot create post", http.StatusInternalServerError)
		return
	}

}
