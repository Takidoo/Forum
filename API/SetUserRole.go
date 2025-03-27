package API

import (
	"Forum/Database"
	"encoding/json"
	"net/http"
	"strconv"
)

func SetUserRole(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		http.Error(w, "Méthode invalide", http.StatusMethodNotAllowed)
		return
	}
	if !Database.UserIsAdmin(w, r) {
		http.Error(w, "You're not admin", http.StatusUnauthorized)
		return
	}
	cookie, _ := r.Cookie("session_id")
	var user User
	resp, _ := http.Get("http://127.0.0.1/UserInfo?session=" + cookie.Value)
	json.NewDecoder(resp.Body).Decode(&user)
	print(user.Role)
	roleInt, converr := strconv.Atoi(r.FormValue("RoleID"))
	if roleInt > user.Role && converr != nil {
		http.Error(w, "Vous ne pouvez pas donner des permissions au dessus des votres", http.StatusUnauthorized)
		return
	}
	_, err := Database.DB.Exec("UPDATE users SET role=? WHERE id=?", r.FormValue("RoleID"), r.FormValue("UserID"))
	if err != nil {
		json.NewEncoder(w).Encode(map[string]string{"error": "Can't set role to user"})
	}
	json.NewEncoder(w).Encode(map[string]string{"success": "Role set to user"})
}
