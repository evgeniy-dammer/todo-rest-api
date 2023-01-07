package service

import (
	"crypto/sha1"
	"fmt"
	todo "github.com/evgeniy-dammer/todo-rest-api"
	"github.com/evgeniy-dammer/todo-rest-api/pkg/repository"
)

const salt = "jdiwbciwcowdkcjndcbdjhb"

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

// generatePasswordHash hashes the password
func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
