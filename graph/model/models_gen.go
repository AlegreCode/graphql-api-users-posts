// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Post struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Text  string `json:"text"`
	User  *User  `json:"user"`
}

type User struct {
	ID       int     `json:"id"`
	Name     string  `json:"name"`
	Lastname string  `json:"lastname"`
	Posts    []*Post `json:"posts"`
}

type InputPost struct {
	Title string `json:"title"`
	Text  string `json:"text"`
}

type InputUser struct {
	Name     string `json:"name"`
	Lastname string `json:"lastname"`
}
