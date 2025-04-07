package Forum

import (
	"Forum/Database"
	"fmt"
)

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Role     int    `json:"role"`
}

func GetUser(session string) (User, error) {
	var user User
	err := Database.DB.QueryRow("SELECT user_id FROM sessions WHERE token = ?", session).Scan(&user.ID)
	if err != nil {

		return User{}, fmt.Errorf("Cannot fetch user infos")
	}
	_ = Database.DB.QueryRow("SELECT username, role FROM users WHERE id = ?", user.ID).Scan(&user.Username, &user.Role)
	return user, nil
}
