package Database

import (
	"crypto/rand"
	"database/sql"
	"encoding/hex"
	"fmt"
	"log"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

var DB *sql.DB

func ConnectDB() {
	var err error
	DB, err = sql.Open("sqlite3", "forum.db")
	if err != nil {
		log.Fatal("Erreur durant la connection a la base de donnée:", err)
	}

	fmt.Println("Connection a la base de donnée avec succés")
	CreateTables()
}

func CreateTables() {
	queries := []string{
		`CREATE TABLE IF NOT EXISTS users (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
            username TEXT UNIQUE NOT NULL,
			password TEXT NOT NULL,
            token TEXT UNIQUE NOT NULL
		);`,
		`CREATE TABLE IF NOT EXISTS threads (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			title TEXT NOT NULL,
			user_id INTEGER NOT NULL,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			FOREIGN KEY(user_id) REFERENCES users(id) ON DELETE CASCADE
		);`,
		`CREATE TABLE IF NOT EXISTS posts (
            id INTEGER PRIMARY KEY AUTOINCREMENT,
			thread_id INTEGER NOT NULL,
			user_id INTEGER NOT NULL,
			content TEXT NOT NULL,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			FOREIGN KEY(thread_id) REFERENCES threads(id) ON DELETE CASCADE,
			FOREIGN KEY(user_id) REFERENCES users(id) ON DELETE CASCADE
        );`,
	}

	for _, query := range queries {
		_, err := DB.Exec(query)
		if err != nil {
			log.Fatal("Erreur durant la création des tables de la base de donnée:", err)
		}
	}
	fmt.Println("Tables de la base de donnée créer avec succes")
}

func CheckIfUserExist(username, password string) bool {
	var exists bool
	query := `SELECT EXISTS (SELECT 1 FROM users WHERE username = ? AND password = ? LIMIT 1);`
	err := DB.QueryRow(query, username, password).Scan(&exists)
	if err != nil {
		if err == sql.ErrNoRows {
			return false
		}
		return false
	}
	return exists
}

func RegisterUser(username, password string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	query := `INSERT INTO users (username, password) VALUES (?, ?)`
	_, err = DB.Exec(query, username, string(hashedPassword))
	return err
}

func LoginUser(username, password string, w http.ResponseWriter) (bool, error) {
	var storedPassword, token string
	var userID int

	query := `SELECT id, password FROM users WHERE username = ?`
	err := DB.QueryRow(query, username).Scan(&userID, &storedPassword)
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

	token = generateToken()

	updateQuery := `UPDATE users SET token = ? WHERE id = ?`
	_, err = DB.Exec(updateQuery, token, userID)
	if err != nil {
		return false, err
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "session_token",
		Value:    token,
		HttpOnly: true,
		Path:     "/",
	})

	return true, nil
}

func generateToken() string {
	bytes := make([]byte, 32)
	_, err := rand.Read(bytes)
	if err != nil {
		log.Fatal("Erreur durant la génération du token:", err)
	}
	return hex.EncodeToString(bytes)
}

func CloseDB() {
	if DB != nil {
		DB.Close()
		fmt.Println("Connection a la base de donnée fermer")
	}
}
