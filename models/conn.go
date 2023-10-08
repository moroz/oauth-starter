package models

import (
	"context"
	"net/http"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/moroz/oauth-starter/config"
)

func ConnectToDB() *sqlx.DB {
	return sqlx.MustConnect("postgres", config.DATABASE_URL)
}

func DBContextMiddleware(db *sqlx.DB) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := context.WithValue(r.Context(), "DB", db)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
