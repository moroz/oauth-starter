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

var GITHUB_CLIENT_ID = MustGetenv("GITHUB_CLIENT_ID")
var GITHUB_CLIENT_SECRET = MustGetenv("GITHUB_CLIENT_SECRET")
var PORT = GetenvWithDefault("PORT", "3000")
var PUBLIC_HOST = MustGetenv("PUBLIC_HOST")
var LISTEN_ON = fmt.Sprintf("0.0.0.0:%s", PORT)
