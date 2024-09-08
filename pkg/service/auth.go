package service

import (
	"crypto/sha512"
	"fmt"
	testmedods "testMEDODS"
	"testMEDODS/pkg/repository"
)

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(user testmedods.User) (int, error) {
	user.Password = generatePasswordHash(user.Password)

	return s.repo.CreateUser(user)
}

const salt = "aksg838tkdnsg"

func generatePasswordHash(password string) string {
	hash := sha512.New()
	hash.Write([]byte(password))
	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
