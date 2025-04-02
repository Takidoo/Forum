package Forum

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Role     int    `json:"role"`
}

type Thread struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
}

type Post struct {
	PostID   int    `json:"post_id"`
	ThreadID int    `json:"thread_id"`
	Owner    int    `json:"owner"`
	Content  string `json:"content"`
	Date     string `json:"created_at"`
}

// Page Data

type HomePageData struct {
	LastedThreads []Thread
}
