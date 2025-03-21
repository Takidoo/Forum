package API

import (
	"Forum/Database"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func RegisterUser(username, password string) error {
	var count int
	checkQuery := `SELECT COUNT(1) FROM users WHERE username = ?;`
	err := Database.DB.QueryRow(checkQuery, username).Scan(&count)
	if err != nil {
		return fmt.Errorf("Erreur lors de la vérification de l'existence de l'utilisateur : %v", err)
	}

	if count > 0 {
		return fmt.Errorf("Nom d'utilisateur déjà pris")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("Erreur lors du hachage du mot de passe : %v", err)
	}
	query := `INSERT INTO users (username, password) VALUES (?, ?)`
	_, err = Database.DB.Exec(query, username, string(hashedPassword))
	if err != nil {
		return fmt.Errorf("Erreur lors de l'insertion du nouvel utilisateur : %v", err)
	}

	return nil
}
