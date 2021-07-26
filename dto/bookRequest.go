package dto

import (
	"net/http"

	"github.com/vindecodex/msgo/errs"
)

type BookRequest struct {
	Id     int    `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Length int    `json:"length"`
}

func (d BookRequest) Validate() *errs.Error {
	if d.Title == "" || d.Author == "" || d.Length == 0 {
		return &errs.Error{
			Code:    http.StatusUnprocessableEntity,
			Message: "Make sure fields are not empty",
		}
	}
	return nil
}
