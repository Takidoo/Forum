package API

import (
	"Forum/Database"
	"encoding/json"
	"net/http"
)

func CreatePost(w http.ResponseWriter, r *http.Request) {
	middleAuth, aerr := Database.MiddlewareAuth(w, r)
	if !middleAuth {
		http.Error(w, aerr.Error(), http.StatusUnauthorized)
		return
	}

	thread_id := r.FormValue("thread_id")
	message := r.FormValue("message")
	if message == "" || thread_id == "" {
		http.Error(w, "Invalid Args", http.StatusBadRequest)
		return
	}

	if !Database.CheckIfThreadExist(thread_id) {
		http.Error(w, "Thread doesn't exist", http.StatusBadRequest)
		return
	}

	cookie, _ := r.Cookie("session_id")
	var user User
	resp, _ := http.Get("http://127.0.0.1/UserInfo?session=" + cookie.Value)
	json.NewDecoder(resp.Body).Decode(&user)
	_, qerr := Database.DB.Exec(`INSERT INTO posts (thread_id, user_id, content) VALUES (?, ?, ?)`, thread_id, user.ID, message)
	if qerr != nil {
		http.Error(w, "Cannot create post", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(map[string]string{"message": "Success"})
}
