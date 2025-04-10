package API

import (
	"Forum/Database"
	"Forum/Forum"
	"encoding/json"
	"net/http"
	"strconv"
)

func SetUserRole(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		http.Error(w, "Méthode invalide", http.StatusMethodNotAllowed)
		return
	}
	session, _ := r.Cookie("session_id")
	if !Forum.UserIsAdmin(session.Value) {
		http.Error(w, "You're not admin", http.StatusUnauthorized)
		return
	}
	var user Forum.User
	user, _ = Forum.GetUser(session.Value)
	roleInt, _ := strconv.Atoi(r.FormValue("RoleID"))
	if roleInt > user.Role {
		http.Error(w, "Vous ne pouvez pas donner des permissions au dessus des votres", http.StatusUnauthorized)
		return
	}
	_, err := Database.DB.Exec("UPDATE users SET role=? WHERE id=?", r.FormValue("RoleID"), r.FormValue("UserID"))
	if err != nil {
		json.NewEncoder(w).Encode(map[string]string{"error": "Can't set role to user"})
		return
	}
	json.NewEncoder(w).Encode(map[string]string{"success": "Role set to user"})
}
