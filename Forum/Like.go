package Forum

import (
	"Forum/Database"
	"fmt"
)

func AddThreadLike(thread_id string, user_id int) error {
	if !CheckIfThreadExist(thread_id) {
		return fmt.Errorf("thread does't exist")
	}
	_, err := Database.DB.Exec("INSERT INTO likes (thread_id, user_id) VALUES (?,?)", thread_id, user_id)
	if err != nil {
		return err
	}
	Database.DB.Exec("UPDATE threads SET likes= likes + 1 WHERE id=?", thread_id)
	return nil
}

func GetThreadLike(thread_id string) int {
	if !CheckIfThreadExist(thread_id) {
		return 0
	}
	var count int
	Database.DB.QueryRow("SELECT COUNT(*) FROM likes WHERE thread_id=?", thread_id).Scan(&count)
	return count
}
