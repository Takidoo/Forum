package Forum

import (
	"Forum/Database"
	"database/sql"
)

type Thread struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
}

func GetLastedThreads(limit int) []Thread {
	rows, _ := Database.DB.Query("SELECT id,title FROM threads WHERE visible=true ORDER BY id DESC LIMIT ?;", limit)
	var Threads []Thread
	for rows.Next() {
		var thread Thread
		rows.Scan(&thread.ID, &thread.Title)
		Threads = append(Threads, thread)
	}
	return Threads
}

func GetMostLikedThreads(limit int) []Thread {
	rows, _ := Database.DB.Query("SELECT id,title FROM threads WHERE visible=true ORDER BY likes ASC LIMIT ?;", limit)
	var Threads []Thread
	for rows.Next() {
		var thread Thread
		rows.Scan(&thread.ID, &thread.Title)
		Threads = append(Threads, thread)
	}
	return Threads
}

func CheckIfThreadExist(thread_id string) bool {
	var title string
	query := `SELECT title FROM threads WHERE id = ?`
	err := Database.DB.QueryRow(query, thread_id).Scan(&title)
	if err != nil {
		if err == sql.ErrNoRows {
			return false
		}
		return false
	}
	return true
}
