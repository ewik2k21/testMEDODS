package repository

import (
	testmedods "testMEDODS"

	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user testmedods.User) (int, error)
	GetUser(email, password string) (testmedods.User, error)
	SetSession(userId int, session testmedods.Session) error
}

type Repository struct {
	Authorization
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
	}
}
