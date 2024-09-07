package service

import "testMEDODS/pkg/repository"

type Authorization interface {
}

type Service struct {
	Authorization
}

func NewService(rep *repository.Repository) *Service {
	return &Service{}
}
