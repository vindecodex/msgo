package controller

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"github.com/vindecodex/msgo/dto"
	"github.com/vindecodex/msgo/mocks/service"
)

func Test_get_all_books_handler_should_return_status_code_200(t *testing.T) {
	// generate controller
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// mockService that implements the service interface
	mockService := service.NewMockBookService(ctrl)

	dummyBooks := []dto.BookResponse{
		{Id: 1, Title: "Test", Author: "Mr.Test", Length: 1000},
		{Id: 2, Title: "Test1", Author: "Mr.Test1", Length: 1000},
	}

	// mock service implements service interface
	mockService.EXPECT().GetAllBooks(gomock.Any()).Return(dummyBooks, nil)

	// Stubs the mock service implementation
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

	if recorder.Code != http.StatusOK {
		t.Error("Fail on testing status code")
	}
}
