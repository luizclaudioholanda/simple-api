package repository

import (
	"api/src/model"

	"github.com/jinzhu/gorm"
)

// Books Represents a book
type Books struct {
	db *gorm.DB
}

// CreateBookRepository Creates a new book repository
func CreateBookRepository(db *gorm.DB) *Books {

	return &Books{db}
}

// CreateBook Create a book
func (repository Books) CreateBook(book model.Book) (uint64, error) {

	// repository.db.AutoMigrate(&book)
	repository.db.Create(&book)
	return uint64(book.Model.ID), nil
}

// FindAllBooks Return all books
func (repository Books) FindAllBooks() ([]model.Book, error) {

	var books []model.Book

	repository.db.Find(&books)
	return books, nil
}

// FindBookByID Return all books
func (repository Books) FindBookByID(bookID int) (model.Book, error) {

	var bookDB model.Book

	repository.db.First(&bookDB, bookID)
	return bookDB, nil
}

// UpdateBookByID Update a book by ID
func (repository Books) UpdateBookByID(bookID int, book model.Book) {

	var bookDB model.Book

	repository.db.First(&bookDB, bookID)
	repository.db.Model(&bookDB).Updates(model.Book{
		Name:     book.Name,
		Author:   book.Author,
		IsbnCode: book.IsbnCode,
	})
}

// DeleteBookByID Delete a book by ID
func (repository Books) DeleteBookByID(bookID int) {

	var bookDB model.Book

	repository.db.First(&bookDB, bookID)
	repository.db.Delete(&bookDB)
}
