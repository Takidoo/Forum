package API

import (
	"Forum/Database"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type Post struct {
	PostID   int    `json:"post_id"`
	ThreadID int    `json:"thread_id"`
	UserID   int    `json:"user_id"`
	Content  string `json:"content"`
	Date     string `json:"created_at"`
}

func FetchThreadPosts(w http.ResponseWriter, r *http.Request) {
	threadID, err := strconv.Atoi(r.FormValue("thread_id"))
	if err != nil {
		http.Error(w, "ID du thread invalide", http.StatusBadRequest)
		return
	}

	cookie, err := r.Cookie("session_id")
	if err != nil {
		http.Error(w, "Token manquant", http.StatusUnauthorized)
		return
	}

	if !Database.UserIsValid(cookie.Value) {
		http.Error(w, "Utilisateur invalide", http.StatusBadRequest)
		return
	}

	rows, err := Database.DB.Query(`SELECT id, thread_id, user_id, content, created_at FROM posts WHERE thread_id = ?`, threadID)
	if err != nil {
		http.Error(w, "Erreur lors de la récupération des messages", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	fmt.Println("Messages récupérés avec succès")

	// Lire les résultats et stocker dans un slice
	var posts []Post
	for rows.Next() {
		var post Post
		if err := rows.Scan(&post.PostID, &post.ThreadID, &post.UserID, &post.Content, &post.Date); err != nil {
			http.Error(w, "Erreur lors de la lecture des données", http.StatusInternalServerError)
			return
		}
		posts = append(posts, post)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(posts)
}
