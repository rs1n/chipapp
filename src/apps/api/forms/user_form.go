package forms

import (
	"log"
	"net/http"
	"strings"

	"github.com/sknv/chip/validate"
	"github.com/sknv/chipapp/src/lib/models"
	"github.com/sknv/chipapp/src/lib/utils"
)

type (
	UserForm struct {
		Login       string          `json:"login" validate:"present,lte=100"`
		Password    string          `json:"password" validate:"present,gte=6"`
		Profile     Profile         `json:"profile"`      // Embeds one profile.
		Images      []*models.Image `json:"images"`       // Embeds many images.
		FollowerIds []string        `json:"follower_ids"` // Has and belongs to many users.
	}

	Profile struct {
		Email  string `json:"email" validate:"omitempty,email,lte=100"`
		Name   string `json:"name" validate:"omitempty,lte=100"`
		Phones string `json:"phones" validate:"omitempty,dive,lte=50"` // TODO: validate phone format
	}
)

func GetUserDupErrorMessage() validate.ValidationErrors {
	return validate.ValidationErrors{
		"UserForm.login": "already exists",
	}
}

func (f *UserForm) FillModel(w http.ResponseWriter, user *models.User) {
	password, err := utils.HashPassword(f.Password)
	if err != nil {
		log.Print("error: ", err)
		utils.RenderStatusAndAbort(w, http.StatusBadRequest)
	}

	user.Login = strings.TrimSpace(f.Login)
	user.Password = password
	user.Profile = models.Profile{
		Email:  f.Profile.Email,
		Name:   f.Profile.Name,
		Phones: f.Profile.Phones,
	}
	user.Images = f.Images
	user.FollowerIds = f.FollowerIds
}
