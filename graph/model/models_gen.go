// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Message struct {
	Message string `json:"Message"`
}

type NewTodo struct {
	Text   string `json:"text"`
	UserID string `json:"userId"`
}

type UpdateTodo struct {
	ID   string `json:"id"`
	Text string `json:"text"`
	Done bool   `json:"done"`
}

type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
