package dto

import (
	"net/http"
	"testing"
)

func Test_user_request_validate_should_return_exact_error_if_empty_fields(t *testing.T) {

	userRequest := UserRequest{
		Username: "Testing",
		// Password: "password",
	}

	userAuthRequest := UserAuthRequest{
		// Username: "Testing",
		Password: "password",
	}

	err1 := userRequest.Validate()
	if err1.Message != "Make sure fields are not empty" {
		t.Error("Invalid error message on testing Validate()")
	}
	if err1.Code != http.StatusUnprocessableEntity {
		t.Error("Invalid status code on testing Validate()")
	}

	err2 := userAuthRequest.Validate()
	if err2.Message != "Make sure fields are not empty" {
		t.Error("Invalid error message on testing Validate()")
	}
	if err2.Code != http.StatusUnprocessableEntity {
		t.Error("Invalid status code on testing Validate()")
	}

}
