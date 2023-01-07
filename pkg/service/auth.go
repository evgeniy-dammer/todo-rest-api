package service

import (
	"crypto/sha1"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	todo "github.com/evgeniy-dammer/todo-rest-api"
	"github.com/evgeniy-dammer/todo-rest-api/pkg/repository"
	"time"
)

const (
	salt       = "jdiwbciwcowdkcjndcbdjhb"
	signingKey = "jsbncvshbvshbvib9872984dcnv"
	tokenTTL   = 12 * time.Hour
)

// tokenClaims
type tokenClaims struct {
	jwt.StandardClaims
	UserId int `json:"user_id"`
}

// AuthService
type AuthService struct {
	repo repository.Authorization
}

// NewAuthService constructor for AuthService
func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

// CreateUser hashes the password and insert User into database
func (s *AuthService) CreateUser(user todo.User) (int, error) {
	user.Password = generatePasswordHash(user.Password)

	return s.repo.CreateUser(user)
}

// GenerateToken generates authorization token
func (s *AuthService) GenerateToken(username string, password string) (string, error) {
	// get user from database
	user, err := s.repo.GetUser(username, generatePasswordHash(password))

	if err != nil {
		return "", err
	}

	// if user exists - create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		user.Id,
	})

	return token.SignedString([]byte(signingKey))
}

// generatePasswordHash hashes the password
func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
