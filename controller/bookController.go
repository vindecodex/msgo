package controller

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/vindecodex/msgo/dto"
	"github.com/vindecodex/msgo/logger"
	"github.com/vindecodex/msgo/service"
)

type BookController struct {
	Service service.BookService
}

func (c BookController) GetAllBooks(w http.ResponseWriter, r *http.Request) {
	logger.Info("GetAllBooks")

	sortBy := r.URL.Query().Get("sort-by")
	if sortBy == "" {
		sortBy = "title"
	}

	books, err := c.Service.GetAllBooks(sortBy)

	if err != nil {
		logger.Error(err.AsMessage())
		writeResponse(w, err.Code, err.AsResponse())
	}

	writeResponse(w, http.StatusOK, books)
}

func (c BookController) GetBook(w http.ResponseWriter, r *http.Request) {
	logger.Info("GetBook")
	idStr := mux.Vars(r)["book_id"]

	id, err := strconv.Atoi(idStr)
	if err != nil {
		logger.Error(err.Error())
	}

	book, e := c.Service.GetBook(id)
	if e != nil {
		logger.Error(e.AsMessage())
		writeResponse(w, e.Code, e.AsResponse())
	}

	writeResponse(w, http.StatusOK, book)
}

func (c BookController) NewBook(w http.ResponseWriter, r *http.Request) {
	logger.Info("NewBook")
	var request dto.BookRequest

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		logger.Error(err.Error())
		writeResponse(w, http.StatusBadRequest, err.Error())
	}

	book, e := c.Service.NewBook(request)
	if e != nil {
		logger.Error(e.AsMessage())
		writeResponse(w, e.Code, e.AsResponse())
	}

	writeResponse(w, http.StatusCreated, book)
}

func (c BookController) UpdateBook(w http.ResponseWriter, r *http.Request) {
	logger.Info("UpdateBook")
	var request dto.BookRequest

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		logger.Error(err.Error())
		writeResponse(w, http.StatusBadRequest, err.Error())
	}

	idStr := mux.Vars(r)["book_id"]

	id, err := strconv.Atoi(idStr)
	if err != nil {
		logger.Error(err.Error())
	}

	request.Id = id

	book, e := c.Service.UpdateBook(request)
	if e != nil {
		logger.Error(e.AsMessage())
		writeResponse(w, e.Code, e.AsResponse())
	}
	writeResponse(w, http.StatusOK, book)
}

func (c BookController) DeleteBook(w http.ResponseWriter, r *http.Request) {
	logger.Info("DeleteBook")

	idStr := mux.Vars(r)["book_id"]

	id, err := strconv.Atoi(idStr)
	if err != nil {
		logger.Error(err.Error())
	}

	e := c.Service.DeleteBook(id)
	if e != nil {
		logger.Error(e.AsMessage())
		writeResponse(w, e.Code, e.AsResponse())
	}

	writeResponse(w, http.StatusOK, nil)

}
