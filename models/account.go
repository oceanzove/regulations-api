package models

import "github.com/golang-jwt/jwt/v5"

type Account struct {
	ID       string `json:"id"`
	Login    string `json:"login"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

type JWTClaims struct {
	ID        string `json:"id"`
	Login     string `json:"login"`
	TokenType string `json:"token_type"`
	jwt.RegisteredClaims
}
