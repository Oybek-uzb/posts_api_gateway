package models

type Post struct {
	ID     int    `json:"id"`
	UserId int    `json:"user_id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

type UpdatePost struct {
	Title  string `json:"title"`
	UserId int    `json:"user_id"`
	Body   string `json:"body"`
}
