package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/moroz/oauth-starter/config"
	"github.com/moroz/oauth-starter/controllers"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", controllers.InitGithubAuth)
	fmt.Printf("Listening on %s\n", config.LISTEN_ON)
	http.ListenAndServe(config.LISTEN_ON, r)
}
