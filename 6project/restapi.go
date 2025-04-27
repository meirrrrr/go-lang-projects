package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"sync"
)

var (
	books     = make([]Book, 0)
	booksLock = sync.Mutex{}
	nextID    = 1
)

type Book struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

func getBooks(w http.ResponseWriter, r *http.Request) {
	booksLock.Lock()
	defer booksLock.Unlock()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

func addBook(w http.ResponseWriter, r *http.Request) {
	booksLock.Lock()
	defer booksLock.Unlock()

	var newBook Book
	if err := json.NewDecoder(r.Body).Decode(&newBook); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	newBook.ID = nextID
	nextID++

	books = append(books, newBook)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newBook)
}

func deleteBook(w http.ResponseWriter, r *http.Request) {
	booksLock.Lock()
	defer booksLock.Unlock()

	idStr := r.URL.Path[len("/books/"):]
	id, err := strconv.Atoi(idStr)
	if err != nil || id <= 0 {
		http.Error(w, "Invalid book ID", http.StatusBadRequest)
		return
	}

	var found bool
	for i, book := range books {
		if book.ID == id {
			books = append(books[:i], books[i+1:]...)
			found = true
			break
		}
	}

	if !found {
		http.Error(w, "Book not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func main() {
	http.HandleFunc("/books", getBooks)              // GET /books
	http.HandleFunc("/books/", deleteBook)           // DELETE /books/{id}
	http.HandleFunc("/books", addBook)               // POST /books

	fmt.Println("Server running on http://localhost:8080/")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Error starting server:", err)
	}
}
