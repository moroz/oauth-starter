package models

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/jmoiron/sqlx"
	"github.com/moroz/oauth-starter/config"
)

func buildClaimsForUser(user User) jwt.Claims {
	now := time.Now()
	exp := now.Add(config.ACCESS_TOKEN_VALIDITY)

	return jwt.RegisteredClaims{
		Subject:   user.Id.String(),
		Issuer:    config.ACCESS_TOKEN_ISSUER,
		IssuedAt:  &jwt.NumericDate{Time: now},
		ExpiresAt: &jwt.NumericDate{Time: exp},
	}
}

func IssueAccessTokenForUser(user User) string {
	claims := buildClaimsForUser(user)
	token := jwt.NewWithClaims(config.ACCESS_TOKEN_SIGNING_METHOD, claims)
	s, _ := token.SignedString(config.ACCESS_TOKEN_SIGNER)
	return s
}

func DecodeAccessToken(tokenString string) (*jwt.RegisteredClaims, error) {
	claims := jwt.RegisteredClaims{}
	_, err := jwt.ParseWithClaims(tokenString, &claims, func(t *jwt.Token) (interface{}, error) {
		return config.ACCESS_TOKEN_SIGNER, nil
	})
	if err != nil {
		return nil, fmt.Errorf("DecodeAccessToken: %w", err)
	}
	return &claims, nil
}

func AuthenticateUserByAccessToken(db *sqlx.DB, tokenString string) (*User, error) {
	claims, err := DecodeAccessToken(tokenString)
	if err != nil {
		fmt.Println(err)
		return nil, errors.New("Invalid or expired JWT")
	}

	user, err := GetUserById(db, claims.Subject)
	if err != nil {
		return nil, fmt.Errorf("AuthenticateUserByAccessToken: %w", err)
	}

	return user, nil
}
