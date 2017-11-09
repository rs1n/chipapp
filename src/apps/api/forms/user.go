package forms

import (
	"net/http"

	"github.com/rs1n/chip/render"

	"github.com/rs1n/chipapp/src/lib/models"
)

type (
	Profile struct {
		Email string `json:"email" validate:"omitempty,email,lte=100"`
	}

	User struct {
		form

		Name        string          `json:"name" validate:"required,lte=100"`
		Profile     Profile         `json:"profile"`      // Embeds one profile.
		Images      []*models.Image `json:"images"`       // Embeds many images.
		FollowerIds []string        `json:"follower_ids"` // Has and belongs to many users.
	}
)

func NewUser(r *http.Request) *User {
	user := &User{}
	render.BindJson(r, user)
	return user
}

func (f *User) FillModel(user *models.User) {
	user.Name = f.Name
	user.Profile = models.Profile{
		Email: f.Profile.Email,
	}
	user.Images = f.Images
	user.FollowerIds = f.FollowerIds
}
