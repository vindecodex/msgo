package controller

import (
	"encoding/json"
	"net/http"

	"github.com/vindecodex/msgo/dto"
	"github.com/vindecodex/msgo/logger"
	"github.com/vindecodex/msgo/service"
)

type UserController struct {
	Service service.UserService
}

func (c UserController) Login(w http.ResponseWriter, r *http.Request) {
	logger.Info("Login")
	var request dto.UserAuthRequest

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		logger.Error(err.Error())
		writeResponse(w, http.StatusBadRequest, err.Error())
	}

	token, e := c.Service.Login(request)
	if e != nil {
		logger.Error(e.AsMessage())
		writeResponse(w, e.Code, e.AsResponse())
	}

	writeResponse(w, http.StatusOK, token)

}

func (c UserController) Register(w http.ResponseWriter, r *http.Request) {
	logger.Info("Register")
	var request dto.UserRequest

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		logger.Error(err.Error())
		writeResponse(w, http.StatusBadRequest, err.Error())
	}

	e := c.Service.Register(request)
	if e != nil {
		logger.Error(e.AsMessage())
		writeResponse(w, e.Code, e.AsResponse())
	}

	writeResponse(w, http.StatusCreated, nil)
}
