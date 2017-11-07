package models

type (
	Image struct {
		Src   string `json:"src"`
		Style string `json:"style"`
	}

	Profile struct {
		Email string `json:"email"`
	}

	User struct {
		Id          string   `json:"id"`
		Name        string   `json:"name"`
		Profile     Profile  `json:"profile"`      // Embeds one profile.
		Images      []*Image `json:"images"`       // Embeds many images.
		FollowerIds []string `json:"follower_ids"` // Has and belongs to many users.

		// Virtual fields.
		Followers  []*User      `json:"followers,omitempty"`  // Has and belongs to many users.
		Microposts []*Micropost `json:"microposts,omitempty"` // Has many microposts.
	}
)
