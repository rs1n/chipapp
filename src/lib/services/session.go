package services

import (
	"context"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/sknv/chip"
	mnware "github.com/sknv/mng/middleware"

	"github.com/sknv/chipapp/src/lib/middleware"
	"github.com/sknv/chipapp/src/lib/models"
	"github.com/sknv/chipapp/src/lib/repositories"
	"github.com/sknv/chipapp/src/lib/utils"
)

type ctxKey string

const (
	ctxKeyCurrentUser = ctxKey("_session.CurrentUser")
	sessionExpiresIn  = 30 * 24 * time.Hour // Expires in 30 days.
)

type (
	Session struct {
		UserRepo *repositories.User `inject:""`
	}

	Authentication struct {
		AuthToken string       `json:"auth_token"`
		User      *models.User `json:"user"`
	}
)

func (_ *Session) CreateAuthentication(
	user *models.User, signingKey []byte,
) *Authentication {
	authToken, err := utils.CreateJwt(signingKey, jwt.MapClaims{
		"exp":     time.Now().Add(sessionExpiresIn).Unix(),
		"user_id": user.Id,
	})
	chip.PanicIfError(err)

	auth := &Authentication{
		AuthToken: authToken,
		User:      user,
	}
	return auth
}

func (s *Session) IsLoggedIn(w http.ResponseWriter, r *http.Request) bool {
	currentUser, _ := s.GetCurrentUser(w, r)
	return currentUser != nil
}

func (s *Session) GetCurrentUser(
	w http.ResponseWriter, r *http.Request,
) (*models.User, *http.Request) {
	currentUser, ok := r.Context().Value(ctxKeyCurrentUser).(*models.User)
	if ok {
		return currentUser, r
	}

	claims, ok := middleware.GetJwtClaims(r)
	if !ok {
		utils.RenderStatusAndAbort(w, http.StatusUnauthorized)
	}

	userId, ok := claims["user_id"].(string)
	if !ok {
		utils.RenderStatusAndAbort(w, http.StatusUnauthorized)
	}

	mgoSession := mnware.GetMgoSession(r)
	currentUser, err := s.UserRepo.FindOneById(mgoSession, userId)
	if err != nil {
		utils.RenderStatusAndAbort(w, http.StatusUnauthorized)
	}

	// Cache current user for request.
	ctx := context.WithValue(r.Context(), ctxKeyCurrentUser, currentUser)
	return currentUser, r.WithContext(ctx)
}
