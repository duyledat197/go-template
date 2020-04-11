package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/duyledat197/go-template/service/user"
	"github.com/duyledat197/go-template/utils"
)

// BEARER ...
var BEARER string = "Bearer "

// Middleware ...
type Middleware func(next http.Handler) http.Handler

// Authentication ...
func Authentication(us user.Service) Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			bearerToken := r.Header.Get("Authorization")
			if strings.HasPrefix(bearerToken, BEARER) {
				http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			}
			token := bearerToken[len(BEARER):len(bearerToken)]
			userID, err := utils.VerifyToken(token)
			if err != nil {
				http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
				return
			}
			user, err2 := us.GetUserByID(userID)
			if err2 != nil {
				http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
				return
			}
			ctx := context.WithValue(r.Context(), "user", &user)
			next.ServeHTTP(w, r.WithContext(ctx))
			return
		})
	}
}
