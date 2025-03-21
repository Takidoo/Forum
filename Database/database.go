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

func CreateTables() {
	queries := []string{
		`CREATE TABLE IF NOT EXISTS users (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
            username TEXT UNIQUE NOT NULL,
			password TEXT NOT NULL,
            token TEXT UNIQUE
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
			log.Fatal("Erreur lors de la création des tables de la base de données:", err)
		}
	}
	fmt.Println("Tables de la base de données créées avec succès")
}

func CheckIfUserExist(username, password string) bool {
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

func RegisterUser(username, password string) error {
	var count int
	checkQuery := `SELECT COUNT(1) FROM users WHERE username = ?;`
	err := DB.QueryRow(checkQuery, username).Scan(&count)
	if err != nil {
		return fmt.Errorf("Erreur lors de la vérification de l'existence de l'utilisateur : %v", err)
	}

	if count > 0 {
		return fmt.Errorf("Nom d'utilisateur déjà pris")
	}
	print("pas existe")

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("Erreur lors du hachage du mot de passe : %v", err)
	}
	print("hask")

	query := `INSERT INTO users (username, password) VALUES (?, ?)`
	_, err = DB.Exec(query, username, string(hashedPassword))
	if err != nil {
		return fmt.Errorf("Erreur lors de l'insertion du nouvel utilisateur : %v", err)
	}

	print("Utilisateur enregistré avec succès !")
	return nil
}

func LoginUser(username, password string, w http.ResponseWriter) (bool, error) {
	var storedPassword, token string
	var userID int

	query := "SELECT id, password FROM users WHERE username = ?"
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
	updateQuery := "UPDATE users SET token = ? WHERE id = ?"
	_, err = DB.Exec(updateQuery, token, userID)
	if err != nil {
		return false, err
	}
	cookie := http.Cookie{
		Name:     "session_id",
		Value:    token,
		Path:     "/",
		HttpOnly: true,
	}
	http.SetCookie(w, &cookie)

	return true, nil
}

func generateToken() string {
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

func CreatePost(threadID, userID int, content string) error {
	query := `INSERT INTO posts (thread_id, user_id, content) VALUES (?, ?, ?)`
	_, err := DB.Exec(query, threadID, userID, content)
	if err != nil {
		return err
	}
	fmt.Println("Post créé avec succès !")
	return nil
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
