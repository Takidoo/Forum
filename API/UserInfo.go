package API

import (
	"Forum/Database"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
}

func UserInfo(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		print("Vision")
		http.Error(w, "Méthode invalide", http.StatusMethodNotAllowed)
		return
	}

	sessionID := r.URL.Query().Get("session")
	if sessionID == "" {
		print("Force")
		http.Error(w, "Invalid session ID", http.StatusBadRequest)
		return
	}

	query := `SELECT id, username FROM users WHERE token = ?`
	row := Database.DB.QueryRow(query, sessionID)

	var user User
	err := row.Scan(&user.ID, &user.Username)
	if err != nil {
		if err == sql.ErrNoRows {
			print("Y'a pas")
			http.Error(w, "User not found", http.StatusNotFound)
			return
		}
		http.Error(w, "Database error", http.StatusInternalServerError)
		log.Println("Erreur SQL:", err)
		return
	}

	fmt.Println("User trouvé:", user.ID, user.Username)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}
