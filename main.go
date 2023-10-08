package main

import (
	"fmt"
	"log"
	"net"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/moroz/oauth-starter/config"
	"github.com/moroz/oauth-starter/controllers"
	"github.com/moroz/oauth-starter/models"
)

func main() {
	db := models.ConnectToDB()

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(models.DBContextMiddleware(db))
	r.Get("/auth/github", controllers.InitGithubAuth)
	r.Get(config.GITHUB_CALLBACK_PATH, controllers.GithubAuthCallback)
	listener, err := net.Listen("tcp", config.LISTEN_ON)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Listening on %s\n", config.LISTEN_ON)
	http.Serve(listener, r)
}
