package API

import (
	"Forum/Database"
	"Forum/Forum"
	"encoding/json"
	"net/http"
)

func CreateCategory(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		http.Error(w, "Méthode invalide", http.StatusMethodNotAllowed)
		return
	}
	session, _ := r.Cookie("session_id")
	if !Forum.UserIsAdmin(session.Value) {
		http.Error(w, "You're not admin", http.StatusUnauthorized)
		return
	}
	_, err := Database.DB.Exec("INSERT INTO categories (name) VALUES (?)", r.FormValue("title"))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"status": "Cannot create a new category"})
		return
	}
	json.NewEncoder(w).Encode(map[string]string{"status": "success"})
}
