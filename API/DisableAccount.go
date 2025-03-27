package API

import (
	"Forum/Database"
	"encoding/json"
	"net/http"
)

func DisableAccount(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		http.Error(w, "Méthode invalide", http.StatusMethodNotAllowed)
		return
	}
	if !Database.UserIsAdmin(w, r) {
		http.Error(w, "You're not admin", http.StatusUnauthorized)
		return
	}
	_, err := Database.DB.Exec("UPDATE users SET account_disabled=? WHERE id=?", r.FormValue("disabled"), r.FormValue("UserID"))
	if err != nil {
		json.NewEncoder(w).Encode(map[string]string{"error": "Can't set role to user"})
		return
	}
	json.NewEncoder(w).Encode(map[string]string{"success": "Account disabled"})
}
