package domain

import (
	"strings"
	"time"

	"github.com/gbrlsnchs/jwt/v3"
	"github.com/vindecodex/msgo/errs"
	"github.com/vindecodex/msgo/logger"
)

type User struct {
	jwt.Payload
	Id       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

var hash = jwt.NewHS256([]byte("theQuickBrownFox"))

type UserRepository interface {
	Authenticate(string, string) (*User, *errs.Error)
	Authorize(string, string) *errs.Error
	NewUser(User) *errs.Error
}

func (d User) Claims() (string, *errs.Error) {
	logger.Info("Claims")
	now := time.Now()

	payload := jwt.Payload{
		Issuer:         "msgo",
		Subject:        "MSGO",
		ExpirationTime: jwt.NumericDate(now.Add(1 * time.Hour)),
		IssuedAt:       jwt.NumericDate(now),
	}

	d.Payload = payload
	token, err := jwt.Sign(d, hash)
	if err != nil {
		return string(""), errs.ServerError("Error on signing token")
	}
	return string(token), nil
}

func (d *User) Verify(token string) *errs.Error {
	_, err := jwt.Verify([]byte(token), hash, &d)
	if err != nil {
		return errs.ServerError("Error on verifying token")
	}
	return nil
}

func (d User) Authorize(url string) *errs.Error {
	if permissions, ok := PERMIT[d.Role]; ok {
		for _, v := range permissions {
			if strings.Contains(url, v) {
				return nil
			}
		}
		return errs.ServerError(d.Role + " is not authorized")
	}

	return errs.ServerError("Role not found")
}
