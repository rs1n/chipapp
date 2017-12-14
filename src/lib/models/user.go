package models

import "upper.io/db.v3/postgresql"

type User struct {
	Base `db:",inline"`

	Email   string           `db:"email" json:"email"`
	Profile postgresql.JSONB `db:"profile" json:"profile"` // Embeds one profile.
	Images  postgresql.JSONB `db:"images" json:"images"`   // Embeds many images.
	// FollowerIds []string `json:"follower_ids"` // Has and belongs to many users.

	// Virtual fields.
	//Followers  []*User      `json:"followers,omitempty"`  // Has and belongs to many users.
	Microposts []*Micropost `json:"microposts,omitempty"` // Has many microposts.
}
