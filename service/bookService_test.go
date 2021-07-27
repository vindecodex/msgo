package service

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/vindecodex/msgo/dto"
	"github.com/vindecodex/msgo/errs"
	"github.com/vindecodex/msgo/mocks/domain"
)

func Test_new_book_should_return_error_on_validation(t *testing.T) {
	request := dto.BookRequest{
		Title:  "",
		Author: "",
		Length: 0,
	}

	bookService := NewDefaultBookService(nil)

	_, err := bookService.NewBook(request)
	if err == nil {
		t.Error("Fail on test validate didn't return error")
	}
}

// Fail Part
func Test_new_book_should_return_error_on_save(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := domain.NewMockBookRepository(ctrl)

	mockRepo.EXPECT().Save(gomock.Any()).Return(nil, errs.ServerError("Unexpected error"))

	bookService := NewDefaultBookService(mockRepo)

	req := dto.BookRequest{
		Title:  "Test",
		Author: "Test",
		Length: 1000,
	}

	_, err := bookService.NewBook(req)
	if err == nil {
		t.Error("Fail on test doesn't return error on save")
	}

}
