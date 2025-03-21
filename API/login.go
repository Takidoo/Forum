package API

import (
	"Forum/Database"
	"encoding/json"
	"fmt"
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Methode non autorisé", http.StatusMethodNotAllowed)
		return
	}

	err := r.ParseForm()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"erreur": "Mauvaise requete"})
		return
	}

	username := r.FormValue("username")
	password := r.FormValue("password")

	w.Header().Set("Content-Type", "application/json")

	if username == "" || password == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"erreur": "Manque le nom d'utilisateur ou mot de passe"})
		return
	}

	success, err := Database.LoginUser(username, password, w)
	if err != nil {
		fmt.Println("Erreur durant le login user:", err)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"erreur": "Erreur server interne"})
		return
	}

	if !success {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(map[string]string{"erreur": "Crédentiels invalide"})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Login fait avec success"})
}
