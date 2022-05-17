package main

import (
	"encoding/json"
	"net/http"
)

type Book struct {
	Title string `json:"Title"`
}

func BooksHandler(w http.ResponseWriter, r *http.Request) {

	books := []Book{
		Book{"The Hobbit"},
		Book{"The Lord of The Rings"},
		Book{"His Dark Materials"},
	}

	json.NewEncoder(w).Encode(books)
}

func main() {

	http.HandleFunc("/books", BooksHandler)
	http.ListenAndServe(":8080", nil)
}
