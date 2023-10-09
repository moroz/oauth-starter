package config

import (
	"encoding/base64"
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func MustGetenv(name string) string {
	value := os.Getenv(name)
	if value == "" {
		msg := fmt.Sprintf("FATAL: Environment variable %s is not set", name)
		panic(msg)
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

func MustGetBase64Env(name string) []byte {
	value := MustGetenv(name)
	raw, err := base64.StdEncoding.DecodeString(value)
	if err != nil {
		msg := fmt.Sprintf("FATAL: Environment variable %s is not a valid base64-encoded string", name)
		panic(msg)
	}
	return raw
}

const COOKIE_PREFIX = "_oauth_starter"
const GITHUB_AUTH_SCOPES = "read:user user:email"
const GITHUB_CALLBACK_PATH = "/auth/github/callback"
const GITHUB_INIT_AUTH_ENDPOINT = "https://github.com/login/oauth/authorize"
const GITHUB_REQUEST_ACCESS_TOKEN_ENDPOINT = "https://github.com/login/oauth/access_token"
const GITHUB_USER_DATA_ENDPOINT = "https://api.github.com/user"
const ACCESS_TOKEN_VALIDITY = 24 * 60 * 60 * time.Second
const ACCESS_TOKEN_COOKIE = COOKIE_PREFIX + "_access_token"
const OAUTH_STATE_COOKIE = COOKIE_PREFIX + "_state"
const USER_CONTEXT_KEY = "USER"

var ACCESS_TOKEN_ISSUER = "oauth-starter"
var ACCESS_TOKEN_SIGNER = MustGetBase64Env("ACCESS_TOKEN_SIGNER")
var DATABASE_URL = MustGetenv("DATABASE_URL")
var GITHUB_CALLBACK_URI = PUBLIC_HOST + GITHUB_CALLBACK_PATH
var GITHUB_CLIENT_ID = MustGetenv("GITHUB_CLIENT_ID")
var GITHUB_CLIENT_SECRET = MustGetenv("GITHUB_CLIENT_SECRET")
var LISTEN_ON = "0.0.0.0:" + PORT
var PORT = GetenvWithDefault("PORT", "3000")
var PUBLIC_HOST = MustGetenv("PUBLIC_HOST")
var ACCESS_TOKEN_SIGNING_METHOD = jwt.SigningMethodHS256
