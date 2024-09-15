package model

type Note struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	UserID  int    `json:"user_id"`
	Content string `json:"content"`
}
