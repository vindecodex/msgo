package controller

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"github.com/vindecodex/msgo/dto"
	"github.com/vindecodex/msgo/errs"
	"github.com/vindecodex/msgo/mocks/service"
)

// This can be refactored but for learning purposes
// this will be leave as it is, so that it can be
// understand how mocks works and how to implement them
// once you understand how it will be easy to refactor
func Test_get_all_books_handler_should_return_status_code_200(t *testing.T) {
	// generate controller
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// mockService that implements the service interface
	mockService := service.NewMockBookService(ctrl)

	// dummy data that we can return
	dummyBooks := []dto.BookResponse{
		{Id: 1, Title: "Test", Author: "Mr.Test", Length: 1000},
		{Id: 2, Title: "Test1", Author: "Mr.Test1", Length: 1000},
	}

	// mock service implements service interface
	mockService.EXPECT().GetAllBooks(gomock.Any()).Return(dummyBooks, nil)

	// Stubs the mock service into the bookController
	bookController := BookController{mockService}

	// creating route
	router := mux.NewRouter()
	router.HandleFunc("/books", bookController.GetAllBooks)

	// requesting the above route
	request, _ := http.NewRequest(http.MethodGet, "/books", nil)

	// recorder is implementation of response writer
	recorder := httptest.NewRecorder()
	// pass the writer and the request
	router.ServeHTTP(recorder, request)

	// assert
	if recorder.Code != http.StatusOK {
		t.Error("Fail on testing status code")
	}
}

func Test_get_all_books_handler_should_return_status_code_500(t *testing.T) {
	// generate controller
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// mockService that implements the service interface
	mockService := service.NewMockBookService(ctrl)

	// mock service implements service interface
	mockService.EXPECT().GetAllBooks(gomock.Any()).Return(nil, errs.ServerError("Unexpected Error"))

	// Stubs the mock service into the bookController
	bookController := BookController{mockService}

	// creating route
	router := mux.NewRouter()
	router.HandleFunc("/books", bookController.GetAllBooks)

	// requesting the above route
	request, _ := http.NewRequest(http.MethodGet, "/books", nil)

	// recorder is implementation of response writer
	recorder := httptest.NewRecorder()
	// pass the writer and the request
	router.ServeHTTP(recorder, request)

	if recorder.Code != http.StatusInternalServerError {
		t.Error("Fail on testing status code")
	}
}
