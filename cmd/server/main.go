package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/re-gis/bms_api/pkg/db"
	"github.com/re-gis/bms_api/pkg/handlers"
	"github.com/re-gis/bms_api/pkg/models"
)

func main() {
	r := mux.NewRouter()
	if err := godotenv.Load(); err != nil {
		log.Println("Warning! No .env file found!")
	}

	// DB CONNECTION
	if err := db.Connect(); err != nil {
		log.Fatalf("Could not initialise database: %s", err)
	}

	db.GetDB().AutoMigrate(&models.Author{}, &models.User{}, &models.Book{}, &models.Loan{})

	defer db.GetDB()

	// ROUTES
	/* BOOK ROUTES */
	r.HandleFunc("/books", handlers.BooksHandler)
	r.HandleFunc("/books/update/{id:[0-9]+}", handlers.BooksHandler).Methods("PUT")
	r.HandleFunc("/books/delete/{id:[0-9]+}", handlers.BooksHandler).Methods("DELETE")

	/* USER ROUTES */
	r.HandleFunc("/users", handlers.UserHandler)

	// SERVER
	// fmt.Println("Server running on port 8080")
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8080", r))
}
