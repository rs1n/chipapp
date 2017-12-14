package forms

import (
	"net/http"

	"github.com/sknv/chip/render"
	"upper.io/db.v3/postgresql"

	"github.com/sknv/chipapp/src/lib/models"
)

type (
	Image struct {
		Src   string `json:"src"`
		Style string `json:"style"`
	}

	Profile struct {
		Name string `json:"name" validate:"omitempty,lte=100"`
	}

	User struct {
		Base

		Email   string   `json:"email" validate:"required,email,lte=100"`
		Profile Profile  `json:"profile"` // Embeds one profile.
		Images  []*Image `json:"images"`  // Embeds many images.
		// FollowerIds []string        `json:"follower_ids"` // Has and belongs to many users.
	}
)

func NewUser(r *http.Request) *User {
	user := &User{}
	render.BindJson(r, user)
	return user
}

func (f *User) FillModel(user *models.User) {
	user.Email = f.Email
	user.Profile = postgresql.JSONB{V: f.Profile}
	user.Images = postgresql.JSONB{V: f.Images}
	// user.FollowerIds = f.FollowerIds
}
