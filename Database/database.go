package Database

import (
	"crypto/rand"
	"database/sql"
	"encoding/hex"
	"fmt"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
	"golang.org/x/crypto/bcrypt"
)

var DB *sql.DB

func ConnectDB() {
	var err error
	DB, err = sql.Open("sqlite3", "forum.db")
	if err != nil {
		log.Fatal("Erreur lors de la connexion à la base de données:", err)
	}

	fmt.Println("Connexion à la base de données réussie")
	CreateTables()
}

func MiddlewareAuth(w http.ResponseWriter, r *http.Request) bool {
	cookie, err := r.Cookie("session_id")
	if err != nil {
		http.Error(w, "Token Needed", http.StatusUnauthorized)
		return false
	}
	var count int
	var account_disabled bool
	err = DB.QueryRow("SELECT COUNT(*), account_disabled FROM users WHERE token = ?", cookie.Value).Scan(&count, &account_disabled)
	if err != nil {
		return false
	} else if account_disabled {
		http.Error(w, "Account is disabled", http.StatusUnauthorized)
		return false
	}
	if count > 0 {
		return true
	} else {
		http.Error(w, "Invalid Session ID", http.StatusBadRequest)
		return false
	}
}

func CreateTables() {
	queries := []string{
		`CREATE TABLE IF NOT EXISTS users (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
            username TEXT UNIQUE NOT NULL,
			password TEXT NOT NULL,
			register TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
            token TEXT UNIQUE,
			account_disabled BOOLEAN DEFAULT false
		);`,
		`CREATE TABLE IF NOT EXISTS threads (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			title TEXT NOT NULL,
			user_id INTEGER NOT NULL,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			visible BOOLEAN DEFAULT true,
			FOREIGN KEY(user_id) REFERENCES users(id) ON DELETE CASCADE
		);`,
		`CREATE TABLE IF NOT EXISTS posts (
            id INTEGER PRIMARY KEY AUTOINCREMENT,
			thread_id INTEGER NOT NULL,
			user_id INTEGER NOT NULL,
			content TEXT NOT NULL,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			visible BOOLEAN DEFAULT true,
			FOREIGN KEY(thread_id) REFERENCES threads(id) ON DELETE CASCADE,
			FOREIGN KEY(user_id) REFERENCES users(id) ON DELETE CASCADE
        );`,
	}

	for _, query := range queries {
		_, err := DB.Exec(query)
		if err != nil {
			log.Fatal("Erreur lors de la création des tables de la base de données:", err)
		}
	}
	fmt.Println("Tables de la base de données créées avec succès")
}

func CheckUserPassword(username, password string) bool {
	var hashedPassword string
	query := `SELECT password FROM users WHERE username = ? LIMIT 1`
	err := DB.QueryRow(query, username).Scan(&hashedPassword)
	if err != nil {
		if err == sql.ErrNoRows {
			return false
		}
		return false
	}

	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		return false
	}

	return true
}

func CheckIfThreadExist(thread_id string) bool {
	var title string
	query := `SELECT title FROM threads WHERE id = ? LIMIT 1`
	err := DB.QueryRow(query, thread_id).Scan(&title)
	if err != nil {
		if err == sql.ErrNoRows {
			return false
		}
		return false
	}
	return true
}

func GenerateToken() string {
	bytes := make([]byte, 32)
	_, err := rand.Read(bytes)
	if err != nil {
		log.Fatal("Erreur lors de la génération du token:", err)
	}
	return hex.EncodeToString(bytes)
}

func CloseDB() {
	if DB != nil {
		DB.Close()
		fmt.Println("Connexion à la base de données fermée")
	}
}

func EditPost(postID, userID int, newContent string) error {
	query := `UPDATE posts SET content = ? WHERE id = ? AND user_id = ?`
	result, err := DB.Exec(query, newContent, postID, userID)
	if err != nil {
		return err
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return fmt.Errorf("Aucun post correspondant trouvé ou l'utilisateur n'est pas propriétaire")
	}

	fmt.Println("Post modifié avec succès !")
	return nil
}

func DeletePost(postID, userID int) error {
	query := `DELETE FROM posts WHERE id = ? AND user_id = ?`
	result, err := DB.Exec(query, postID, userID)
	if err != nil {
		return err
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return fmt.Errorf("Aucun post correspondant trouvé ou l'utilisateur n'est pas propriétaire")
	}

	fmt.Println("Post supprimé avec succès !")
	return nil
}
