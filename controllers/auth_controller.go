package controllers

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"log"
	"net/http"

	"github.com/moroz/oauth-starter/config"
	"github.com/moroz/oauth-starter/models"
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

func setAccessTokenCookie(w http.ResponseWriter, cookie string) {
	http.SetCookie(w, &http.Cookie{
		Name:     config.ACCESS_TOKEN_COOKIE,
		Value:    cookie,
		Secure:   true,
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
		Path:     "/",
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
	w.Header().Set("Cache-Control", "no-store")
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
	resp, err := oauth.RequestGithubUserData(tokenResp.AccessToken)

	if err != nil {
		log.Printf("GithubAuthCallback: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Internal server error")
		return
	}

	db := models.GetDBFromRequest(r)
	user, err := models.GetOrCreateUserByEmail(db, resp.Email)

	if err != nil {
		log.Printf("GithubAuthCallback: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Internal server error")
		return
	}

	token := models.IssueAccessTokenForUser(*user)
	setAccessTokenCookie(w, token)
	fmt.Fprintln(w, resp)
}
