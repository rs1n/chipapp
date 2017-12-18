package models

type (
	Image struct {
		Src   string `json:"src"`
		Style string `json:"style"`
	}

	Profile struct {
		Name string `json:"name"`
	}

	User struct {
		Base `bson:",inline"`

		Email       string   `json:"email"`
		Profile     Profile  `json:"profile"`                          // Embeds one profile.
		Images      []*Image `json:"images"`                           // Embeds many images.
		FollowerIds []string `bson:"follower_ids" json:"follower_ids"` // Has and belongs to many users.

		// Virtual fields.
		Followers  []*User      `bson:"-" json:"followers,omitempty"`  // Has and belongs to many users.
		Microposts []*Micropost `bson:"-" json:"microposts,omitempty"` // Has many microposts.
	}
)
