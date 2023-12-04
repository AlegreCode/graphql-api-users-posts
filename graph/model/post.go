package model

type Post struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Text   string `json:"text"`
	User   *User  `json:"user"`
	UserID int    `json:"userId"`
}
