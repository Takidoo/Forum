package API

import (
	"Forum/Database"
	"encoding/json"
	"net/http"
)

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Role     string `json:"role"`
}

func UserInfo(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Méthode invalide", http.StatusMethodNotAllowed)
		return
	}

	sessionID := r.URL.Query().Get("session")
	if sessionID == "" {
		http.Error(w, "Invalid session ID", http.StatusBadRequest)
		return
	}
	var user User
	err := Database.DB.QueryRow("SELECT user_id FROM sessions WHERE token = ?", sessionID).Scan(&user.ID)
	if err != nil {
		http.Error(w, "Invalid Session ID", http.StatusBadRequest)
		return
	}
	_ = Database.DB.QueryRow("SELECT username, role FROM users WHERE id = ?", user.ID).Scan(&user.Username, &user.Role)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}
