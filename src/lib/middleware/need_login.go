package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/sknv/chip/render"

	"github.com/sknv/chipapp/src/lib/utils"
)

type ctxKey string

const ctxKeyJwtClaims = ctxKey("_jwt.Claims")

func NeedLogin(signingKey []byte) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			tokenString := tokenFromHeader(r)
			if tokenString == "" {
				render.Status(w, http.StatusUnauthorized)
				return
			}

			claims, err := utils.ParseJwt(tokenString, signingKey)
			if err != nil {
				render.Status(w, http.StatusUnauthorized)
				return
			}

			ctx := context.WithValue(r.Context(), ctxKeyJwtClaims, claims)
			next.ServeHTTP(w, r.WithContext(ctx))
		}

		return http.HandlerFunc(fn)
	}
}

func GetJwtClaims(r *http.Request) (jwt.MapClaims, bool) {
	claims, ok := r.Context().Value(ctxKeyJwtClaims).(jwt.MapClaims)
	return claims, ok
}

// tokenFromHeader tries to retrieve the token string from the
// "Authorization" request header: "Authorization: BEARER T".
func tokenFromHeader(r *http.Request) string {
	// Get token from authorization header.
	bearer := r.Header.Get("Authorization")
	if len(bearer) > 7 && strings.ToUpper(bearer[0:6]) == "BEARER" {
		return bearer[7:]
	}
	return ""
}
