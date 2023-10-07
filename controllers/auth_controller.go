package controllers

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"net/http"

	"github.com/moroz/oauth-starter/config"
	"github.com/moroz/oauth-starter/oauth"
)

func generateSecretBytes(length int) string {
	buf := make([]byte, length)
	rand.Read(buf)
	return hex.EncodeToString(buf)
}

func getCookieByName(cookies []*http.Cookie, name string) string {
	for _, cookie := range cookies {
		if cookie.Name == name {
			return cookie.Value
		}
	}
	return ""
}

func setOAuthStateCookie(w http.ResponseWriter, state string) {
	http.SetCookie(w, &http.Cookie{
		Name:     config.OAUTH_STATE_COOKIE,
		Value:    state,
		Secure:   true,
		HttpOnly: true,
		SameSite: http.SameSiteNoneMode,
	})
}

func InitGithubAuth(w http.ResponseWriter, r *http.Request) {
	state := generateSecretBytes(8)
	url, err := oauth.BuildGithubRedirectURL(state)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Internal server error")
		return
	}

	setOAuthStateCookie(w, state)
	http.Redirect(w, r, url, http.StatusMovedPermanently)
}

func GithubAuthCallback(w http.ResponseWriter, r *http.Request) {
	code := r.URL.Query().Get("code")
	if code == "" {
		w.WriteHeader(http.StatusUnprocessableEntity)
		fmt.Fprintf(w, "Unprocessable entity")
		return
	}

	state := r.URL.Query().Get("state")
	stateCookie := getCookieByName(r.Cookies(), config.OAUTH_STATE_COOKIE)

	if state != stateCookie {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprintf(w, "Invalid state cookie")
		return
	}

	tokenResp, _ := oauth.RequestGithubAccessToken(code)
	resp, _ := oauth.RequestGithubUserData(tokenResp.AccessToken)
	fmt.Fprintln(w, resp)
}
