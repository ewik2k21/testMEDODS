package repository

import (
	"fmt"
	testmedods "testMEDODS"

	"github.com/jmoiron/sqlx"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (a *AuthPostgres) CreateUser(user testmedods.User) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (username, email, password_hash) values ($1, $2, $3) RETURNING id", usersTable)
	row := a.db.QueryRow(query, user.Username, user.Email, user.Password)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (a *AuthPostgres) GetUser(email, password string) (testmedods.User, error) {
	var user testmedods.User
	query := fmt.Sprintf("SELECT id FROM %s WHERE email=$1 AND password_hash=$2", usersTable)
	err := a.db.Get(&user, query, email, password)

	return user, err
}

func (a *AuthPostgres) SetSession(userId int, session testmedods.Session) error {
	return nil
}
