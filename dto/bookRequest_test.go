package dto

import (
	"net/http"
	"testing"
)

func Test_book_request_validate_should_return_exact_error_if_empty_fields(t *testing.T) {

	request := BookRequest{
		Title:  "Testing",
		Author: "MrTest",
		// Length: 1000,
	}

	err := request.Validate()
	if err.Message != "Make sure fields are not empty" {
		t.Error("Invalid error message on testing Validate()")
	}
	if err.Code != http.StatusUnprocessableEntity {
		t.Error("Invalid status code on testing Validate()")
	}
}
