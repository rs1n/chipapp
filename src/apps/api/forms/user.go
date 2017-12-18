package forms

import (
	"net/http"

	"github.com/sknv/chipapp/src/lib/models"
)

type (
	Profile struct {
		Name string `json:"name" validate:"omitempty,lte=100"`
	}

	User struct {
		Base

		Email       string          `json:"email" validate:"required,email,lte=100"`
		Profile     Profile         `json:"profile"`      // Embeds one profile.
		Images      []*models.Image `json:"images"`       // Embeds many images.
		FollowerIds []string        `json:"follower_ids"` // Has and belongs to many users.
	}
)

func NewUser(r *http.Request) (*User, error) {
	user := &User{}
	err := bindJson(r, user)
	return user, err
}

func (f *User) FillModel(user *models.User) {
	user.Email = f.Email
	user.Profile = models.Profile{
		Name: f.Profile.Name,
	}
	user.Images = f.Images
	user.FollowerIds = f.FollowerIds
}
