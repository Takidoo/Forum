package API

import (
	"Forum/Database"
	"Forum/Forum"
	"encoding/json"
	"net/http"
)

func DisableAccount(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		http.Error(w, "Méthode invalide", http.StatusMethodNotAllowed)
		return
	}
	session, _ := r.Cookie("session_id")
	if !Forum.UserIsAdmin(session.Value) {
		http.Error(w, "You're not admin", http.StatusUnauthorized)
		return
	}
	disable := r.FormValue("disabled") == "0"
	_, err := Database.DB.Exec("UPDATE users SET account_disabled=? WHERE id=?", disable, r.FormValue("UserID"))
	if err != nil {
		json.NewEncoder(w).Encode(map[string]string{"error": "Can't set role to user"})
		return
	}
	json.NewEncoder(w).Encode(map[string]string{"success": "Account disabled"})
}
