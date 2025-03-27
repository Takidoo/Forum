package API

import (
	"Forum/Database"
	"encoding/json"
	"net/http"
)

func CreateThread(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Méthode Invalide", http.StatusBadRequest)
		return
	}
	if r.FormValue("title") == "" {
		http.Error(w, "Titre Invalide", http.StatusBadRequest)
		return
	}
	middleAuth, aerr := Database.MiddlewareAuth(w, r)
	if !middleAuth {
		http.Error(w, aerr.Error(), http.StatusUnauthorized)
		return
	}
	cookie, _ := r.Cookie("session_id")
	resp, err := http.Get("http://127.0.0.1/UserInfo?session=" + cookie.Value)
	if err != nil {
		http.Error(w, "Impossible de récupérer les infos de l'utilisateur", http.StatusInternalServerError)
		return
	}
	var user User
	json.NewDecoder(resp.Body).Decode(&user)
	query := `INSERT INTO threads (title, user_id) VALUES (?, ?)`
	_, qerr := Database.DB.Exec(query, r.FormValue("title"), user.ID)
	if qerr != nil {
		http.Error(w, "Impossible de créer le thread", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(map[string]string{"message": "Success"})
}
