package API

import (
	"Forum/Database"
	"database/sql"
	"encoding/json"
	"net/http"
)

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
}

func GetUserInfo(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Méthode Invalide", http.StatusBadRequest)
		return
	}
	if !Database.MiddlewareAuth(w, r) {
		return
	}
	cookie, err := r.Cookie("session_id")
	query := `SELECT id, username FROM users WHERE token = ?`
	row := Database.DB.QueryRow(query, cookie.Value)

	var user User
	err = row.Scan(&user.ID, &user.Username)
	if err != nil {
		if err == sql.ErrNoRows {
			return
		}
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}
