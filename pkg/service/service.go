package service

import (
	testmedods "testMEDODS"
	"testMEDODS/pkg/repository"
)

type Authorization interface {
	CreateUser(user testmedods.User) (int, error)
	GenerateToken(email, password string) (string, error)
	ParseToken(token string) (int, error)
	NewRefreshToken() (string, error)
}

type Service struct {
	Authorization
}

func NewService(rep *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(rep.Authorization),
	}
}
