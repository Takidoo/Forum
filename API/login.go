package API

import (
	"Forum/Database"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

func Login(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Méthode non autorisée", http.StatusMethodNotAllowed)
		return
	}

	err := r.ParseForm()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"erreur": "Mauvaise requête"})
		return
	}

	username := r.FormValue("username")
	password := r.FormValue("password")

	w.Header().Set("Content-Type", "application/json")

	if username == "" || password == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"erreur": "Nom d'utilisateur ou mot de passe manquant"})
		return
	}

	success, err := LoginUser(username, password, w)
	if err != nil {
		fmt.Println("Erreur durant le login user:", err)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"erreur": "Erreur interne du serveur"})
		return
	}

	if !success {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(map[string]string{"erreur": "Identifiants invalides"})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Connexion réussie avec succès"})
}

func LoginUser(username, password string, w http.ResponseWriter) (bool, error) {
	var storedPassword, token string
	var userID int

	query := "SELECT id, password FROM users WHERE username = ? AND account_disabled=false"
	err := Database.DB.QueryRow(query, username).Scan(&userID, &storedPassword)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}
		return false, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(storedPassword), []byte(password))
	if err != nil {
		return false, nil
	}

	token = Database.GenerateToken()
	query = "INSERT INTO sessions (token, user_id) VALUES (?,?)"
	_, err = Database.DB.Exec(query, token, userID)
	if err != nil {
		return false, err
	}
	cookie := http.Cookie{
		Name:     "session_id",
		Value:    token,
		Path:     "/",
		HttpOnly: true,
	}
	http.SetCookie(w, &cookie)

	return true, nil
}
