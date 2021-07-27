package domain

import (
	"github.com/vindecodex/msgo/dto"
	"github.com/vindecodex/msgo/errs"
)

type Book struct {
	Id     int    `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Length int    `json:"length"`
}

//go:generate mockgen -destination=../mocks/domain/mockBookRepository.go -package=domain github.com/vindecodex/msgo/domain BookRepository
type BookRepository interface {
	GetAll(string) ([]Book, *errs.Error)
	GetById(int) (*Book, *errs.Error)
	Save(Book) (*Book, *errs.Error)
	Update(Book) (*Book, *errs.Error)
	Delete(int) *errs.Error
}

func (b Book) ToResponseDto() *dto.BookResponse {
	return &dto.BookResponse{
		Id:     b.Id,
		Title:  b.Title,
		Author: b.Author,
		Length: b.Length,
	}
}
