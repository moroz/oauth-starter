package controllers

import (
	"fmt"
	"net/http"

	"github.com/moroz/oauth-starter/config"
)

func InitGithubAuth(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, config.GITHUB_CLIENT_ID)
}
