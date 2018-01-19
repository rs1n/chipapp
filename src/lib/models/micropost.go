package models

type Micropost struct {
	Base `bson:",inline"`

	Content string `json:"content"`
	UserId  string `bson:"user_id" json:"user_id"` // Belongs to a user.

	// Virtual fields.
	User *User `bson:"-" json:"user,omitempty"` // Belongs to a user.
}
