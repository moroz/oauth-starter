package controllers

import (
	"context"
	"net/http"

	"github.com/jmoiron/sqlx"
	"github.com/moroz/oauth-starter/config"
	"github.com/moroz/oauth-starter/models"
)

func UserMiddleware(db *sqlx.DB) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			authCookie := getCookieByName(r.Cookies(), config.ACCESS_TOKEN_COOKIE)
			user, _ := models.AuthenticateUserByAccessToken(db, authCookie)
			ctx := context.WithValue(r.Context(), config.USER_CONTEXT_KEY, user)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func GetUserFromRequest(r *http.Request) *models.User {
	return r.Context().Value(config.USER_CONTEXT_KEY).(*models.User)
}
