package API

import (
	"Forum/Database"
	"encoding/json"
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	password := r.FormValue("password")
	w.Header().Set("Content-Type", "application/json")
	if username != "" && password != "" {
		if !Database.CheckIfUserExist(username, password) {
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode("{Error : \"Nom d'utilisateur ou mot de passe invalide\"}")
		} else {
			w.WriteHeader(http.StatusOK)
			Database.LoginUser(username, password, w)
		}
	} else {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode("{Error : \"Invalid Args\"}")
	}
}
