package domain

import (
	"github.com/jmoiron/sqlx"
	"github.com/vindecodex/msgo/errs"
	"github.com/vindecodex/msgo/logger"
)

type UserRepositoryAdapter struct {
	client *sqlx.DB
}

func NewUserRepositoryAdapter(client *sqlx.DB) UserRepositoryAdapter {
	return UserRepositoryAdapter{client}
}

func (d UserRepositoryAdapter) Authenticate(username string, password string) (*User, *errs.Error) {
	logger.Info("Authenticate")

	q := "SELECT username, role FROM users WHERE username=? AND password=?"

	user := User{}

	err := d.client.Get(&user, q, username, password)
	if err != nil {
		logger.Error(err.Error())
		return nil, errs.UnAuthorizedError("invalid credentials")
	}

	return &user, nil
}

func (d UserRepositoryAdapter) Authorize(token string, url string) *errs.Error {
	logger.Info("Authorize")
	user := User{}

	err := user.Verify(token)
	if err != nil {
		return err
	}

	err = user.Authorize(url)
	if err != nil {
		return err
	}

	return nil
}

func (d UserRepositoryAdapter) NewUser(user User) *errs.Error {
	logger.Info("Register")

	q := "INSERT INTO users(username, password, role) VALUES(?, ?, ?)"

	_, err := d.client.Exec(q, user.Username, user.Password, user.Role)
	if err != nil {
		logger.Error(err.Error())
		return errs.ServerError("Unexpected DB error")
	}

	return nil
}
