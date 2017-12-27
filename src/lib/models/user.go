package models

import "github.com/sknv/chipapp/src/lib/utils"

type (
	User struct {
		Base `bson:",inline"`

		Login       string   `json:"login"`
		Password    string   `json:"-"`
		Profile     Profile  `json:"profile"`                          // Embeds one profile.
		Images      []*Image `json:"images"`                           // Embeds many images.
		FollowerIds []string `bson:"follower_ids" json:"follower_ids"` // Has and belongs to many users.

		// Virtual fields.
		Followers  []*User      `bson:"-" json:"followers,omitempty"`  // Has and belongs to many users.
		Microposts []*Micropost `bson:"-" json:"microposts,omitempty"` // Has many microposts.
	}

	Image struct {
		Src   string `json:"src"`
		Style string `json:"style"`
	}

	Profile struct {
		Email  string `json:"email"`
		Name   string `json:"name"`
		Phones string `json:"phones"`
	}
)

func (u *User) Authenticate(password string) bool {
	return utils.CheckPasswordHash(password, u.Password)
}
