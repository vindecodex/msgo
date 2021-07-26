package domain

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/vindecodex/msgo/errs"
	"github.com/vindecodex/msgo/logger"
)

type BookRepositoryAdapter struct {
	client *sqlx.DB
}

func NewBookRepositoryAdapter(client *sqlx.DB) BookRepositoryAdapter {
	return BookRepositoryAdapter{client}
}

func (b BookRepositoryAdapter) GetAll(sortBy string) ([]Book, *errs.Error) {
	logger.Info("GetAll")
	sortBy = "ORDER BY " + sortBy

	q := "SELECT id, title, author, length FROM books " + sortBy

	books := []Book{}

	err := b.client.Select(&books, q)
	if err != nil {
		logger.Error(err.Error())
		return nil, errs.ServerError("Unexpected DB error")
	}

	return books, nil
}

func (b BookRepositoryAdapter) GetById(id int) (*Book, *errs.Error) {
	logger.Info("GetById")
	q := "SELECT id, title, author, length FROM books WHERE id = ?"

	var book Book

	err := b.client.Get(&book, q, id)
	if err != nil {
		if err == sql.ErrNoRows {
			logger.Error(err.Error())
			return nil, errs.NotFoundError("No book found")
		}
		logger.Error(err.Error())
		return nil, errs.ServerError("Unexpected DB error")
	}
	return &book, nil
}

func (b BookRepositoryAdapter) Save(book Book) (*Book, *errs.Error) {
	logger.Info("Save")
	q := "INSERT INTO books(title, author, length) VALUES(?, ?, ?)"

	res, err := b.client.Exec(q, book.Title, book.Author, book.Length)
	if err != nil {
		logger.Error(err.Error())
		return nil, errs.ServerError("Unexpected DB error")
	}

	id, err := res.LastInsertId()
	if err != nil {
		logger.Error(err.Error())
		return nil, errs.ServerError("Unexpected error")
	}

	book.Id = int(id)

	return &book, nil

}

func (b BookRepositoryAdapter) Update(book Book) (*Book, *errs.Error) {
	logger.Info("Update")
	q := "UPDATE books SET title=?, author=?, length=? WHERE id=?"

	_, err := b.client.Exec(q, book.Title, book.Author, book.Length, book.Id)
	if err != nil {
		logger.Error(err.Error())
		return nil, errs.ServerError("Unexpected DB error")
	}

	return &book, nil
}

func (b BookRepositoryAdapter) Delete(id int) *errs.Error {
	logger.Info("Delete")
	q := "DELETE FROM books WHERE id=?"

	_, err := b.client.Exec(q, id)
	if err != nil {
		logger.Error(err.Error())
		return errs.ServerError("Unexpected DB error")
	}
	return nil
}
