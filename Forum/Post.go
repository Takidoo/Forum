package Forum

type Post struct {
	PostID   int    `json:"post_id"`
	ThreadID int    `json:"thread_id"`
	Owner    int    `json:"owner"`
	Content  string `json:"content"`
	Date     string `json:"created_at"`
}
