package controllers

import (
	"fmt"
	"net/http"
)

func UserInfo(w http.ResponseWriter, r *http.Request) {
	user := GetUserFromRequest(r)

	fmt.Fprintln(w, user)
}
