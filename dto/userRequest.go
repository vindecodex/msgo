package dto

import (
	"net/http"

	"github.com/vindecodex/msgo/errs"
)

type UserRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

type UserAuthRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (d UserRequest) Validate() *errs.Error {
	if d.Username == "" || d.Password == "" || d.Role == "" {
		return &errs.Error{
			Code:    http.StatusUnprocessableEntity,
			Message: "Make sure fields are not empty",
		}
	}
	return nil
}

func (d UserAuthRequest) Validate() *errs.Error {
	if d.Username == "" || d.Password == "" {
		return &errs.Error{
			Code:    http.StatusUnprocessableEntity,
			Message: "Make sure fields are not empty",
		}
	}
	return nil
}
