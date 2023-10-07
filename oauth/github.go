package oauth

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/google/go-querystring/query"
	"github.com/moroz/oauth-starter/config"
)

type githubOAuthInitParams struct {
	State       string `url:"state"`
	ClientId    string `url:"client_id"`
	Scope       string `url:"scope"`
	RedirectURI string `url:"redirect_uri"`
}

func BuildGithubRedirectURL(state string) (string, error) {
	params := githubOAuthInitParams{
		State:       state,
		ClientId:    config.GITHUB_CLIENT_ID,
		Scope:       config.GITHUB_AUTH_SCOPES,
		RedirectURI: config.GITHUB_CALLBACK_URI,
	}
	values, err := query.Values(params)
	if err != nil {
		return "", err
	}

	qs := values.Encode()
	return fmt.Sprintf("%s?%s", config.GITHUB_INIT_AUTH_ENDPOINT, qs), nil
}

type githubOAuthAccessTokenParams struct {
	ClientID     string `url:"client_id"`
	ClientSecret string `url:"client_secret"`
	Code         string `url:"code"`
}

type GithubOAuthAccessTokenResponse struct {
	AccessToken string `json:"access_token"`
	Scope       string `json:"scope"`
	TokenType   string `json:"token_type"`
}

func RequestGithubAccessToken(oauthCode string) (*GithubOAuthAccessTokenResponse, error) {
	params := githubOAuthAccessTokenParams{
		ClientID:     config.GITHUB_CLIENT_ID,
		ClientSecret: config.GITHUB_CLIENT_SECRET,
		Code:         oauthCode,
	}
	values, err := query.Values(params)
	if err != nil {
		return nil, err
	}
	body := strings.NewReader(values.Encode())

	client := &http.Client{}
	r, _ := http.NewRequest("POST", config.GITHUB_REQUEST_ACCESS_TOKEN_ENDPOINT, body)
	r.Header.Add("Accept", "application/json")
	resp, err := client.Do(r)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result GithubOAuthAccessTokenResponse
	err = json.Unmarshal(respBody, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

type GithubUserData struct {
	Login     string `json:"login"`
	AvatarURL string `json:"avatar_url"`
	Email     string `json:"email"`
}

func RequestGithubUserData(bearerToken string) (*GithubUserData, error) {
	client := &http.Client{}
	r, _ := http.NewRequest("GET", config.GITHUB_USER_DATA_ENDPOINT, nil)
	r.Header.Add("Authorization", fmt.Sprintf("Bearer %s", bearerToken))
	resp, err := client.Do(r)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result GithubUserData
	err = json.Unmarshal(respBody, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
