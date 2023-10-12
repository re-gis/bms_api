package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/re-gis/bms_api/pkg/db"
	"github.com/re-gis/bms_api/pkg/models"
)

func BooksHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		getBooks(w, r)
	case http.MethodPost:
		addBook(w, r)

	case http.MethodPut:
		updateBook(w, r)
	case http.MethodDelete:
		deleteBook(w, r)

	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("Method not allowed!"))
	}
}

func addBook(w http.ResponseWriter, r *http.Request) {
	var book models.Book

	err := json.NewDecoder(r.Body).Decode(&book)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Failed to decode book data!"))
		return
	}
	if book.AuthorId == 0 || book.Genre == "" || book.ISBN == "" || book.Title == "" {
		http.Error(w, "All credentials are require!", http.StatusBadRequest)
		return
	}

	result := db.GetDB().Create(&book)
	if result.Error != nil {
		http.Error(w, "Failed to create the book!", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(book)
}

func getBooks(w http.ResponseWriter, r *http.Request) {
	var books []models.Book

	// get books from database
	result := db.GetDB().Find(&books)

	if result.Error != nil {
		http.Error(w, "Failed to retrieve books...", http.StatusInternalServerError)
		return
	}

	jsonResponse, err := json.Marshal(books)
	if err != nil {
		http.Error(w, "Failed to parse the books...", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}

func updateBook(w http.ResponseWriter, r *http.Request) {
	// get id from url
	parts := strings.Split(r.URL.Path, "/")
	IdStr := parts[len(parts)-1]

	if IdStr == "" {
		http.Error(w, "Id required!", http.StatusBadRequest)
		return
	}

	Id, err := strconv.Atoi(IdStr)

	if err != nil {
		http.Error(w, "Invalid Id provided!", http.StatusBadRequest)
		return
	}

	// decode the body
	var updateBook models.Book
	err = json.NewDecoder(r.Body).Decode(&updateBook)
	if err != nil {
		http.Error(w, "Failed to decode the body", http.StatusInternalServerError)
		return
	}

	// get the book from db
	var book models.Book

	if err = db.GetDB().First(&book, Id).Error; err != nil {
		http.Error(w, "Book not found!", http.StatusNotFound)
		return
	}

	if updateBook.Title != "" {
		book.Title = updateBook.Title
	}
	if updateBook.ISBN != "" {
		book.ISBN = updateBook.ISBN
	}

	if updateBook.AuthorID != 0 {
		book.AuthorId = updateBook.AuthorId
	}

	if updateBook.Genre != "" {
		book.Genre = updateBook.Genre
	}

	if err = db.GetDB().Save(&book).Error; err != nil {
		http.Error(w, "Error while updating the book...", http.StatusInternalServerError)
		return
	}

	jsonResponse, err := json.Marshal(book)
	if err != nil {
		http.Error(w, "Error while parsing the book!", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Book updated successfully"))
	w.Write(jsonResponse)
}

func deleteBook(w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(r.URL.Path, "/")
	IdStr := parts[len(parts)-1]

	if IdStr == "" {
		http.Error(w, "Id required!", http.StatusBadRequest)
		return
	}

	Id, err := strconv.Atoi(IdStr)
	if err != nil {
		http.Error(w, "An error occurred while formating the id", http.StatusInternalServerError)
		return
	}
	// get the book from database
	var book models.Book

	if err := db.GetDB().First(&book, Id).Error; err != nil {
		http.Error(w, "Book not found!", http.StatusNotFound)
		return
	}

	// delete it
	if err = db.GetDB().Delete(&book).Error; err != nil {
		http.Error(w, "Error while deleting the book...", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Successfully deleted the book!"))

}
