package Database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
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
            user_id INTEGER NOT NULL
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

func CloseDB() {
	if DB != nil {
		DB.Close()
		fmt.Println("Connection a la base de donnée fermer")
	}
}
