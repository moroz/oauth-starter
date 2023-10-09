package controllers

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/jmoiron/sqlx"
	"github.com/moroz/oauth-starter/config"
	"github.com/moroz/oauth-starter/models"
)

func Routes(db *sqlx.DB) http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(models.DBContextMiddleware(db))
	r.Use(UserMiddleware(db))

	r.Get("/", UserInfo)
	r.Get("/auth/github", InitGithubAuth)
	r.Get(config.GITHUB_CALLBACK_PATH, GithubAuthCallback)
	return r
}
