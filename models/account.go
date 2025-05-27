package models

import "github.com/golang-jwt/jwt/v5"

type Account struct {
	ID       string `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

type JWTClaims struct {
	ID        string `json:"id"`
	Email     string `json:"email"`
	TokenType string `json:"token_type"`
	jwt.RegisteredClaims
}
