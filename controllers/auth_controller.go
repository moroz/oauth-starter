package controllers

import (
	"fmt"
	"net/http"
)

func InitGithubAuth(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Test")
}
