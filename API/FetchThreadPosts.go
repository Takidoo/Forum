package API

import (
	"Forum/Database"
	"Forum/Forum"
	"encoding/json"
	"fmt"
	"net/http"
)

func FetchThreadPosts(w http.ResponseWriter, r *http.Request) {
	middleAuth, aerr := Database.MiddlewareAuth(w, r)
	if !middleAuth {
		http.Error(w, aerr.Error(), http.StatusUnauthorized)
		return
	}
	if !Forum.CheckIfThreadExist(r.FormValue("thread_id")) {
		http.Error(w, "Invalid Thread ID", http.StatusBadRequest)
		return
	}

	rows, err := Database.DB.Query(`SELECT id, thread_id, user_id, content, created_at FROM posts WHERE thread_id = ? AND visible=true`, r.FormValue("thread_id"))
	if err != nil {
		http.Error(w, "Erreur lors de la récupération des messages", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	fmt.Println("Messages récupérés avec succès")

	var posts []Forum.Post
	for rows.Next() {
		var post Forum.Post
		if err := rows.Scan(&post.PostID, &post.ThreadID, &post.Owner, &post.Content, &post.Date); err != nil {
			http.Error(w, "Erreur lors de la lecture des données", http.StatusInternalServerError)
			return
		}
		posts = append(posts, post)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(posts)
}
