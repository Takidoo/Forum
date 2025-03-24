package API

import (
	"Forum/Database"
	"encoding/json"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

func RegisterUser(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	password := r.FormValue("password")
	if username == "" || password == "" {
		http.Error(w, "Invalid Agrs", http.StatusBadRequest)
	}
	w.Header().Set("Content-Type", "application/json")
	var count int
	checkQuery := `SELECT COUNT(1) FROM users WHERE username = ?;`
	err := Database.DB.QueryRow(checkQuery, username).Scan(&count)
	if err != nil {
		json.NewEncoder(w).Encode(map[string]string{"message": "Can't verify if user exist"})
		return
	}

	if count > 0 {
		json.NewEncoder(w).Encode(map[string]string{"message": "Username taken"})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		json.NewEncoder(w).Encode(map[string]string{"message": "Can't create user"})
		http.Error(w, "Can't create user", http.StatusInternalServerError)
		return
	}
	query := `INSERT INTO users (username, password) VALUES (?, ?)`
	_, err = Database.DB.Exec(query, username, string(hashedPassword))
	if err != nil {
		json.NewEncoder(w).Encode(map[string]string{"message": "Can't create user"})
		http.Error(w, "Can't create user", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(map[string]string{"message": "Success"})
}
