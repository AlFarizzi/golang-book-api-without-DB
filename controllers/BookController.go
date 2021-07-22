package Controller

import (
	"encoding/json"
	"math/rand"
	"net/http"
	. "simple-rest-api-mux-golang/models"
	"strconv"

	"github.com/gorilla/mux"
)

var books []Book

func Init() {
	books = append(books, Book{"B0001", "ISBN-0001", "Seni Untuk Bersikap Bodo Amat", Author{"Mark", "Ornsen"}})
	books = append(books, Book{"B0002", "ISBN-0002", "Seni Untuk Mempengaruhi Orang Lain", Author{"Mark", "Ornsen"}})
}

func GetBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

func GetBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)["id"]
	for _, book := range books {
		if book.ID == params {
			json.NewEncoder(w).Encode(book)
		}
	}
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var filteredBooks []Book
	id := mux.Vars(r)["id"]
	for _, book := range books {
		if book.ID != id {
			filteredBooks = append(filteredBooks, book)
		}
	}
	books = filteredBooks
	json.NewEncoder(w).Encode(books)
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var book Book
	err := json.NewDecoder(r.Body).Decode(&book)
	if err == nil {
		book.ID = strconv.Itoa(rand.Intn(100000))
		books = append(books, book)
		json.NewEncoder(w).Encode(books)
	}
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	var book Book
	id := mux.Vars(r)["id"]
	err := json.NewDecoder(r.Body).Decode(&book)
	if err == nil {
		for index, dataBook := range books {
			if dataBook.ID == id {
				dataBook.Isbn = book.Isbn
				dataBook.Title = book.Title
				dataBook.Author.Firstname = book.Author.Firstname
				dataBook.Author.Lastname = book.Author.Lastname
				books[index] = dataBook
				break
			}
		}
		json.NewEncoder(w).Encode(books)
	}
}
