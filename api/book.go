package api

import "net/http"
import "encoding/json"

// book type withn Name, Author, and ISBN
type Book struct {
    Title       string `json:"title"`
    Author      string `json:"author"`
    ISBN        string `json:"isbn"`
    Description string `json:"description,omitempty"`
}

var books = map[string]Book{
    "0345391802": Book{Title: "The Hitchhiker's Guide to the Galaxy", Author: "Douglas Adams", ISBN: "0345391802"},
    "0000000000": Book{Title: "Cloud Native Go", Author: "M.-Leander Reimer", ISBN: "0000000000"},
}

// AllBooks returns a slice of all books
func AllBooks() []Book {
    values := make([]Book, len(books))
    idx := 0
    for _, book := range books {
        values[idx] = book
        idx++
    }
    return values
}

// ToJSON on
func (b Book) ToJSON() []byte {
    book, err := json.Marshal(b)
    if err != nil {
        panic(err)
    }

    return book
}

// FromJSON on
func FromJSON(data []byte) Book {
    book := Book{}
    err := json.Unmarshal(data, &book)
    if err != nil {
        panic(err)
    }

    return book
}

// BooksHandleFunc gets all books
func BooksHandleFunc(w http.ResponseWriter, r *http.Request) {
    b, err := json.Marshal(books)
    if err != nil {
        panic(err)
    }

    w.Header().Add("Content-Type", "text/json")
    w.Write(b)
}

// BooksHandleFunc a book
func BoosHandleFunc(w http.ResponseWriter, r *http.Request) {
    b, err := json.Marshal(books)
    if err != nil {
        panic(err)
    }

    w.Header().Add("Content-Type", "text/json")
    w.Write(b)
}
