package repository

import (
	testmedods "testMEDODS"

	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user testmedods.User) (int, error)
}

type Repository struct {
	Authorization
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
	}
}
