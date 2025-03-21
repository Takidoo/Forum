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
	}
	if !Database.MiddlewareAuth(w, r) {
		return
	}
	resp, err := http.Get("/GetUserInfo")
	if err != nil {
		http.Error(w, "Impossible de récupérer les infos de l'utilisateur", http.StatusInternalServerError)
	}
	var user User
	json.NewDecoder(resp.Body).Decode(&user)
	query := `INSERT INTO thread (title, user_id) VALUES (?, ?)`
	_, qerr := Database.DB.Exec(query, r.FormValue("title"), user.ID)
	if qerr != nil {
		http.Error(w, "Impossible de créer le thread", http.StatusInternalServerError)
		return
	}
	print("Thread créer avec succés par " + user.Username)
}
