package Forum

import (
	"Forum/Database"
)

func GetLastedThreads(limit int) []Thread {
	rows, _ := Database.DB.Query("SELECT id,title FROM threads ORDER BY id DESC LIMIT ?;", limit)
	var Threads []Thread
	for rows.Next() {
		var thread Thread
		rows.Scan(&thread.ID, &thread.Title)
		Threads = append(Threads, thread)
	}
	return Threads
}
