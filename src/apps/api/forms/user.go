package forms

import (
	"net/http"

	"github.com/rs1n/chip/render"

	"github.com/rs1n/chipapp/src/lib/models"
)

type User struct {
	Name        string          `json:"name"`
	Profile     models.Profile  `json:"profile"`      // Embeds one profile.
	Images      []*models.Image `json:"images"`       // Embeds many images.
	FollowerIds []string        `json:"follower_ids"` // Has and belongs to many users.
}

func NewUser(r *http.Request) *User {
	user := &User{}
	render.BindJson(r, user)
	return user
}

func (f *User) FillModel(user *models.User) {
	user.Name = f.Name
	user.Profile = f.Profile
	user.Images = f.Images
	user.FollowerIds = f.FollowerIds
}
