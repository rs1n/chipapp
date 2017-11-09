package models

type Micropost struct {
	model

	Id      string `json:"id"`
	Content string `json:"content"`
	UserId  string `json:"user_id"` // Belongs to user.

	// Virtual fields.
	User *User `json:"user,omitempty"` // Belongs to user.
}
