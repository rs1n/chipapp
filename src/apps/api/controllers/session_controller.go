package controllers

import (
	"log"
	"net/http"
	"strings"

	"github.com/sknv/chip/render"
	mware "github.com/sknv/mng/middleware"

	"github.com/sknv/chipapp/src/apps/api/forms"
	"github.com/sknv/chipapp/src/lib/repositories"
	"github.com/sknv/chipapp/src/lib/services"
	"github.com/sknv/chipapp/src/lib/utils"
)

type SessionController struct {
	*BaseController

	SessionService *services.SessionService
	UserRepo       *repositories.UserRepository
}

func NewSessionController() *SessionController {
	return &SessionController{
		BaseController: NewBaseController(),
		SessionService: services.NewSessionService(nil),
		UserRepo:       repositories.NewUserRepository(),
	}
}

func (c *SessionController) Login(w http.ResponseWriter, r *http.Request) {
	// BindJson the request body to a form and validate it.
	sessionForm := c.bindRequestToSessionForm(w, r)

	mgoSession := mware.GetMgoSession(r)
	login := strings.TrimSpace(sessionForm.Login)
	user, err := c.UserRepo.FindOneByLogin(mgoSession, login)
	if err != nil {
		log.Print("error: ", err)
		utils.RenderStatusAndAbort(w, http.StatusUnauthorized)
	}

	if !user.Authenticate(sessionForm.Password) {
		log.Print("warn: user is not authenticated")
		utils.RenderStatusAndAbort(w, http.StatusUnauthorized)
	}

	// Create authentication info.
	auth := c.SessionService.CreateAuthentication(user, c.Config.SecretKey)
	render.Json(w, http.StatusCreated, auth)
}

func (c *SessionController) bindRequestToSessionForm(
	w http.ResponseWriter, r *http.Request,
) *forms.SessionForm {
	sessionForm := &forms.SessionForm{}
	utils.BindRequest(w, r, sessionForm)
	return sessionForm
}
