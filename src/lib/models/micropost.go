package models

type Micropost struct {
	Base `db:",inline"`

	Content string `db:"content" json:"content"`
	UserId  string `db:"user_id" json:"user_id"` // Belongs to user.

	// Virtual fields.
	User *User `json:"user,omitempty"` // Belongs to user.
}
