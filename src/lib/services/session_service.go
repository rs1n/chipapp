package services

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/sknv/chip"
	"github.com/sknv/mng/odm/document"

	"github.com/sknv/chipapp/src/lib/middleware"
	"github.com/sknv/chipapp/src/lib/utils"
)

type ctxKey string

const (
	ctxKeyCurrentUser = ctxKey("_session.CurrentUser")
	sessionExpiresIn  = 30 * 24 * time.Hour // Expires in 30 days.
)

type (
	SessionService struct {
		UserFinder IUserFinderById
	}

	Authentication struct {
		AuthToken string               `json:"auth_token"`
		User      document.IIdentifier `json:"user"`
	}

	IUserFinderById interface {
		FindOneById(*http.Request, string) (interface{}, error)
	}
)

func NewSessionService(userFinder IUserFinderById) *SessionService {
	return &SessionService{userFinder}
}

func (_ *SessionService) CreateAuthentication(
	user document.IIdentifier, signingKey []byte,
) *Authentication {
	authToken, err := utils.CreateJwt(signingKey, jwt.MapClaims{
		"exp":     time.Now().Add(sessionExpiresIn).Unix(),
		"user_id": user.GetId(),
	})
	chip.PanicIfError(err)

	auth := &Authentication{
		AuthToken: authToken,
		User:      user,
	}
	return auth
}

func (s *SessionService) GetCurrentUser(
	w http.ResponseWriter, r *http.Request,
) (interface{}, *http.Request) {
	currentUser := r.Context().Value(ctxKeyCurrentUser)
	if currentUser != nil {
		return currentUser, r
	}

	claims, ok := middleware.GetJwtClaims(r)
	if !ok {
		log.Print("info: jwt claims do not exist")
		return nil, r
	}

	userId, ok := claims["user_id"].(string)
	if !ok || userId == "" {
		log.Print("error: user id does not exist in jwt claims")
		return nil, r
	}

	currentUser, err := s.UserFinder.FindOneById(r, userId)
	if err != nil {
		log.Print("error: ", err)
		return nil, r
	}

	// Cache current user for request.
	ctx := context.WithValue(r.Context(), ctxKeyCurrentUser, currentUser)
	return currentUser, r.WithContext(ctx)
}
