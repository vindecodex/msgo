package service

import (
	"github.com/vindecodex/msgo/domain"
	"github.com/vindecodex/msgo/dto"
	"github.com/vindecodex/msgo/errs"
)

type BookService interface {
	GetAllBooks(string) ([]dto.BookResponse, *errs.Error)
	GetBook(int) (*dto.BookResponse, *errs.Error)
	NewBook(dto.BookRequest) (*dto.BookResponse, *errs.Error)
	UpdateBook(dto.BookRequest) (*dto.BookResponse, *errs.Error)
	DeleteBook(int) *errs.Error
}

type DefaultBookService struct {
	repo domain.BookRepository
}

func NewDefaultBookService(repo domain.BookRepository) DefaultBookService {
	return DefaultBookService{repo}
}

func (s DefaultBookService) GetAllBooks(sortBy string) ([]dto.BookResponse, *errs.Error) {
	books, err := s.repo.GetAll(sortBy)
	if err != nil {
		return nil, err
	}

	response := []dto.BookResponse{}
	for _, b := range books {
		book := dto.BookResponse{
			Id:     b.Id,
			Title:  b.Title,
			Author: b.Author,
			Length: b.Length,
		}
		response = append(response, book)
	}
	return response, nil
}

func (s DefaultBookService) GetBook(id int) (*dto.BookResponse, *errs.Error) {
	b, err := s.repo.GetById(id)
	if err != nil {
		return nil, err
	}

	return &dto.BookResponse{
		Id:     b.Id,
		Title:  b.Title,
		Author: b.Author,
		Length: b.Length,
	}, nil
}

func (s DefaultBookService) NewBook(b dto.BookRequest) (*dto.BookResponse, *errs.Error) {
	e := b.Validate()
	if e != nil {
		return nil, e
	}

	book := domain.Book{
		Title:  b.Title,
		Author: b.Author,
		Length: b.Length,
	}

	result, err := s.repo.Save(book)
	if err != nil {
		return nil, err
	}

	bookResp := result.ToResponseDto()
	return bookResp, nil
}

func (s DefaultBookService) UpdateBook(b dto.BookRequest) (*dto.BookResponse, *errs.Error) {

	oldBook, err := s.repo.GetById(b.Id)
	if err != nil {
		return nil, err
	}

	if b.Title == "" {
		b.Title = oldBook.Title
	}

	if b.Author == "" {
		b.Author = oldBook.Author
	}

	if b.Length == 0 {
		b.Length = oldBook.Length
	}

	e := b.Validate()
	if e != nil {
		return nil, e
	}

	book := domain.Book{
		Id:     b.Id,
		Title:  b.Title,
		Author: b.Author,
		Length: b.Length,
	}

	result, err := s.repo.Update(book)
	if err != nil {
		return nil, err
	}

	bookResp := result.ToResponseDto()

	return bookResp, nil
}

func (s DefaultBookService) DeleteBook(id int) *errs.Error {
	return s.repo.Delete(id)
}
