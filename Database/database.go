package database

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
            active INT NOT NULL,
            token TEXT NOT NULL
		);`,
    queries := []string{
        `CREATE TABLE IF NOT EXISTS posts (
            id INTEGER PRIMARY KEY AUTOINCREMENT,
            title TEXT UNIQUE NOT NULL,
            email TEXT UNIQUE NOT NULL,
            password TEXT NOT NULL
        );`,
    }

func CheckIfUserExist(username string, password string) bool {

}