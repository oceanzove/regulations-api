package service

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"regulations-api/models"
	"strings"
	"time"
)

const (
	errorEmptyAuthHeader = "authorization header is empty"
	errorGetUserContext  = "failed to get user context"
	errorService         = "service error"
)

const (
	AccessTokenTTL  = time.Minute * 60
	RefreshTokenTTL = time.Hour * 24 * 7
)

type JWTTokenService struct {
	config models.ServerConfig
}

func NewJWTTokenService(config models.ServerConfig) *JWTTokenService {
	return &JWTTokenService{config: config}
}

func getSecretKey(config models.ServerConfig) []byte {
	if config.JWTSecretKey == "" {
		logrus.Fatalf("JWT секретный ключ не установлен в конфигурации")
	}
	return []byte(config.JWTSecretKey)
}

func (s *JWTTokenService) GenerateAccessToken(email string) (string, error) {
	claims := models.JWTClaims{
		Email:     email,
		TokenType: "access",
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(AccessTokenTTL).UTC()),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ID:        uuid.New().String(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(getSecretKey(s.config))
}

func (s *JWTTokenService) GenerateRefreshToken(email string) (string, error) {
	claims := models.JWTClaims{
		Email:     email,
		TokenType: "refresh",
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(RefreshTokenTTL).UTC()),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ID:        uuid.New().String(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(getSecretKey(s.config))
}

func (s *JWTTokenService) GenerateAccessFromRefresh(refreshToken string) (string, error) {
	if strings.HasPrefix(refreshToken, "Bearer ") {
		refreshToken = strings.TrimPrefix(refreshToken, "Bearer ")
		refreshToken = strings.TrimSpace(refreshToken)
		logrus.Infof("Stripped 'Bearer ' from token, now: %v", refreshToken)
	}

	claims, err := s.ParseToken(refreshToken)
	if err != nil {
		logrus.Error("Failed to parse refresh token:", err)
		return "", err
	}

	if claims.TokenType != "refresh" {
		logrus.Error("Received token is not of type 'refresh', token type: ", claims.TokenType)
		return "", errors.New("incorrect token type")
	}

	if time.Now().UTC().After(claims.ExpiresAt.Time) {
		logrus.Error("Refresh token is expired.")
		return "", errors.New("refresh token is expired")
	}

	accessToken, err := s.GenerateAccessToken(claims.Email)
	if err != nil {
		logrus.Error("Error generating access token:", err)
		return "", err
	}

	return accessToken, nil
}

func (s *JWTTokenService) ParseToken(tokenString string) (*models.JWTClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &models.JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		return getSecretKey(s.config), nil
	})
	if err != nil || !token.Valid {
		return nil, err
	}

	claims, ok := token.Claims.(*models.JWTClaims)
	if !ok {
		return nil, errors.New("не удалось извлечь данные из токена")
	}

	return claims, nil
}

func (s *JWTTokenService) RefreshToken(token string) (string, error) {
	claims, err := s.ParseToken(token)
	if err != nil {
		logrus.Error(err.Error())
		return "", err
	}
	if claims.ExpiresAt.Time.Before(time.Now()) {
		return "", errors.New("refresh токен истёк, требуется повторная аутентификация")
	}

	newRefreshToken, err := s.GenerateRefreshToken(claims.Email)
	if err != nil {
		logrus.Error(err.Error())
		return "", err
	}

	return newRefreshToken, nil
}
