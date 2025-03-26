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

func MiddlewareAuth(w http.ResponseWriter, r *http.Request) (bool, error) {
	cookie, err := r.Cookie("session_id")
	if err != nil {
		return false, fmt.Errorf("Token Needed")
	}
	var user_id int
	var account_disabled bool
	err = DB.QueryRow("SELECT user_id FROM sessions WHERE token = ?", cookie.Value).Scan(&user_id)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, fmt.Errorf("Invalid ID")
		}
		return false, fmt.Errorf("Invalid ID")
	}
	_ = DB.QueryRow("SELECT account_disabled FROM users WHERE id = ?", user_id).Scan(&account_disabled)
	if account_disabled {
		return false, fmt.Errorf("Account is disabled, please contact support")
	}
	return true, nil
}

func CreateTables() {
	queries := []string{
		`CREATE TABLE IF NOT EXISTS users (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
            username TEXT UNIQUE NOT NULL,
			password TEXT NOT NULL,
			register TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
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
		`
		CREATE TABLE IF NOT EXISTS sessions (
			token TEXT NOT NULL,
			user_id INT NOT NULL,
			PRIMARY KEY (token),
			FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
	);
		`,
		`
			CREATE UNIQUE INDEX IF NOT EXISTS idx_token ON sessions(token);
		`,
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
