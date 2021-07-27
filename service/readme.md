Mock integration for testing routes

This can be implemented in the command line but this is a good way of doing it. So that we will know that we are generating mocks on this interface.

```go
//go:generate mockgen -destination=../mocks/service/mockBookService.go -package=service github.com/vindecodex/msgo/service BookService
type BookService interface {
	GetAllBooks(string) ([]dto.BookResponse, *errs.Error)
	GetBook(int) (*dto.BookResponse, *errs.Error)
	NewBook(dto.BookRequest) (*dto.BookResponse, *errs.Error)
	UpdateBook(dto.BookRequest) (*dto.BookResponse, *errs.Error)
	DeleteBook(int) *errs.Error
}
```

`go:generate mockgen` runs mockgen tool to generate mocks for BookService interface

Running the mockgen to generate mocks and start testing:
`go generate ./...`

This command will search the entire codebase directory that contains comments to generate the mock

---
The implementation for testing the route will be found inside the controller directory that contains _test.go

For the sake of demo purposes Not all routes have tests. Its very tideous to write mock test

more information using [golang/mock](https://github.com/golang/mock)


