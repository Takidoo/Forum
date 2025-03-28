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
	if r.FormValue("title") == "" || r.FormValue("category") == "" {
		http.Error(w, "Invalid Args", http.StatusBadRequest)
		return
	}
	if !Database.CheckIfCategoryExist(r.FormValue("category")) {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"status": "Invalid Category"})
		return
	}
	middleAuth, aerr := Database.MiddlewareAuth(w, r)
	if !middleAuth {
		http.Error(w, aerr.Error(), http.StatusUnauthorized)
		return
	}
	cookie, _ := r.Cookie("session_id")
	resp, err := http.Get("http://127.0.0.1/api/UserInfo?session=" + cookie.Value)
	if err != nil {
		http.Error(w, "Impossible de récupérer les infos de l'utilisateur", http.StatusInternalServerError)
		return
	}
	var user User
	json.NewDecoder(resp.Body).Decode(&user)
	query := `INSERT INTO threads (title, user_id, category) VALUES (?, ?, ?)`
	_, qerr := Database.DB.Exec(query, r.FormValue("title"), user.ID, r.FormValue("category"))
	if qerr != nil {
		print(qerr.Error())
		http.Error(w, "Impossible de créer le thread", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(map[string]string{"message": "Success"})
}
