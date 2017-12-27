package controllers

import (
	"net/http"
	"strings"

	"github.com/sknv/chip/render"
	mnware "github.com/sknv/mng/middleware"

	"github.com/sknv/chipapp/src/apps/api/forms"
	"github.com/sknv/chipapp/src/lib/repositories"
	"github.com/sknv/chipapp/src/lib/services"
	"github.com/sknv/chipapp/src/lib/utils"
)

type Session struct {
	*Base `inject:""`

	SessionService *services.Session  `inject:""`
	UserRepo       *repositories.User `inject:""`
}

func (c *Session) Login(w http.ResponseWriter, r *http.Request) {
	// BindJson the request body to a form and validate it.
	sessionForm := c.bindRequestToSessionForm(w, r)

	mgoSession := mnware.GetMgoSession(r)
	login := strings.TrimSpace(sessionForm.Login)
	user, err := c.UserRepo.FindOneByLogin(mgoSession, login)
	if err != nil || !user.Authenticate(sessionForm.Password) {
		utils.RenderStatusAndAbort(w, http.StatusUnauthorized)
	}

	// Create authentication info.
	auth := c.SessionService.CreateAuthentication(user, c.Config.SecretKey)
	render.Json(w, http.StatusCreated, auth)
}

func (c *Session) bindRequestToSessionForm(w http.ResponseWriter, r *http.Request) *forms.SessionForm {
	sessionForm := &forms.SessionForm{}
	utils.BindRequest(w, r, sessionForm)
	return sessionForm
}
