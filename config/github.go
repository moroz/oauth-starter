package config

import (
	"fmt"
	"log"
	"os"
)

func MustGetenv(name string) string {
	value := os.Getenv(name)
	if value == "" {
		log.Fatalf("FATAL: Environment variable %s is not set\n", name)
	}
	return value
}

func GetenvWithDefault(name, defaultValue string) string {
	value := os.Getenv(name)
	if value == "" {
		return defaultValue
	}
	return value
}

const COOKIE_PREFIX = "_oauth_starter"
const GITHUB_CALLBACK_PATH = "/auth/github/callback"
const GITHUB_AUTH_SCOPES = "read:user user:email"
const GITHUB_INIT_AUTH_ENDPOINT = "https://github.com/login/oauth/authorize"
const GITHUB_REQUEST_ACCESS_TOKEN_ENDPOINT = "https://github.com/login/oauth/access_token"
const GITHUB_USER_DATA_ENDPOINT = "https://api.github.com/user"

var GITHUB_CLIENT_ID = MustGetenv("GITHUB_CLIENT_ID")
var GITHUB_CLIENT_SECRET = MustGetenv("GITHUB_CLIENT_SECRET")
var PORT = GetenvWithDefault("PORT", "3000")
var PUBLIC_HOST = MustGetenv("PUBLIC_HOST")
var LISTEN_ON = fmt.Sprintf("0.0.0.0:%s", PORT)
var GITHUB_CALLBACK_URI = fmt.Sprintf("%s%s", PUBLIC_HOST, GITHUB_CALLBACK_PATH)
var OAUTH_STATE_COOKIE = fmt.Sprintf("%s_state", COOKIE_PREFIX)
