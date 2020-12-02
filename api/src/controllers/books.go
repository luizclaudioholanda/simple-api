package controllers

import (
	"api/src/database"
	"api/src/model"
	"api/src/repository"
	"api/src/response"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// CreateBook Create a book
func CreateBook(w http.ResponseWriter, r *http.Request) {

	request, err := ioutil.ReadAll(r.Body)

	if err != nil {
		response.Error(w, http.StatusInternalServerError, err)
		return
	}

	var book model.Book
	if err = json.Unmarshal(request, &book); err != nil {
		response.Error(w, http.StatusBadRequest, err)
		return
	}

	dbConn, err := database.Connect()
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err)
		return
	}

	bookRepo := repository.CreateBookRepository(dbConn)

	bookID, err := bookRepo.CreateBook(book)

	if err != nil {
		response.Error(w, http.StatusInternalServerError, err)
		return
	}

	response.JSON(w, http.StatusOK, bookID)
}

// FindAllBooks Find all books
func FindAllBooks(w http.ResponseWriter, r *http.Request) {

	dbConn, err := database.Connect()
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err)
		return
	}

	bookRepo := repository.CreateBookRepository(dbConn)

	book, err := bookRepo.FindAllBooks()

	response.JSON(w, http.StatusOK, book)
}

// FindBook Find a book by ID
func FindBook(w http.ResponseWriter, r *http.Request) {

	dbConn, err := database.Connect()
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err)
		return
	}

	params := mux.Vars(r)

	bookRepo := repository.CreateBookRepository(dbConn)

	bookID, err := strconv.Atoi(params["bookId"])
	if err != nil {
		response.Error(w, http.StatusBadRequest, err)
		return
	}

	book, err := bookRepo.FindBookByID(bookID)

	response.JSON(w, http.StatusOK, book)
}

// UpdateBook Update a book by ID
func UpdateBook(w http.ResponseWriter, r *http.Request) {

	request, err := ioutil.ReadAll(r.Body)

	if err != nil {
		response.Error(w, http.StatusInternalServerError, err)
		return
	}

	var book model.Book
	if err = json.Unmarshal(request, &book); err != nil {
		response.Error(w, http.StatusBadRequest, err)
		return
	}

	dbConn, err := database.Connect()
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err)
		return
	}

	params := mux.Vars(r)

	bookRepo := repository.CreateBookRepository(dbConn)

	bookID, err := strconv.Atoi(params["bookId"])
	if err != nil {
		response.Error(w, http.StatusBadRequest, err)
		return
	}

	bookRepo.UpdateBookByID(bookID, book)

	response.JSON(w, http.StatusOK, nil)
}

// DeleteBook Delete a book by ID
func DeleteBook(w http.ResponseWriter, r *http.Request) {
	dbConn, err := database.Connect()
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err)
		return
	}

	params := mux.Vars(r)

	bookRepo := repository.CreateBookRepository(dbConn)

	bookID, err := strconv.Atoi(params["bookId"])
	if err != nil {
		response.Error(w, http.StatusBadRequest, err)
		return
	}

	bookRepo.DeleteBookByID(bookID)

	response.JSON(w, http.StatusOK, nil)
}
