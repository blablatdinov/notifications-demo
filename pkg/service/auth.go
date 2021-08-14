package service

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"time"

	"github.com/blablatdinov/notifications-demo"
	"github.com/blablatdinov/notifications-demo/pkg/repository"

	"github.com/dgrijalva/jwt-go"
)

type AuthService struct {
	repo repository.Authorization
}

type tokenClaims struct {
	jwt.StandardClaims
	UserId   int    `json:"user_id"`
	Username string `json:"username"`
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(user notifications.User) (int, error) {
	user.Password = generatePasswordHash(user.Password)
	return s.repo.CreateUser(user)
}

const (
	salt       = "someString"
	signingKey = "qrkjk#4#%35FSFJlja#4353KSFjH"
	tokenTTL   = 12 * time.Hour
)

func (s *AuthService) ParseToken(accessToken string) (int, string, error) {
	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}
		return []byte(signingKey), nil
	})
	if err != nil {
		return 0, "", err
	}
	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return 0, "", errors.New("token claims are not of type *tokenClaims")
	}
	return claims.UserId, claims.Username, nil
}

func (s *AuthService) GetUsers() ([]notifications.User, error) {
	users, err := s.repo.GetUsers()
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (s *AuthService) GenerateToken(user notifications.User) (string, error) {
	username := user.Username
	user, err := s.repo.GetUser(user.Username, generatePasswordHash(user.Password))
	if err != nil {
		return "", err
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		user.Id,
		username,
	})
	return token.SignedString([]byte(signingKey))
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
