package model

type Todo struct {
	ID   string `json:"id" bson:"_id"`
	Text string `json:"text"`
	Done bool   `json:"done"`
	User *User  `json:"user"`
}
