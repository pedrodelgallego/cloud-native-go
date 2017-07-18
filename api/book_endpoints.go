package api

import (
    "encoding/json"
    "io/ioutil"
    "net/http"
)

// ToJSON to be used for marshalling of Book type
func (b Book) ToJSON() []byte {
    ToJSON, err := json.Marshal(b)
    if err != nil {
        panic(err)
    }
    return ToJSON
}


// BooksHandleFunc to be used as http.HandleFunc for Book API
func BooksHandleFunc(w http.ResponseWriter, r *http.Request) {
    switch method := r.Method; method {
    case http.MethodGet:
        w = getBookHandle(w)
    case http.MethodPost:
        w = updateBookHandle(r, w)
    default:
        w.WriteHeader(http.StatusBadRequest)
        w.Write([]byte("Unsupported request method."))
    }
}

func updateBookHandle(r *http.Request, w http.ResponseWriter) http.ResponseWriter {
    body, err := ioutil.ReadAll(r.Body)
    if hasInternalServerError(err, w) {
        return w
    }
    
    book := FromJSON(body)
    isbn, created := CreateBook(book)
    if created {
        w.Header().Add("Location", "/api/books/"+isbn)
        w.WriteHeader(http.StatusCreated)
    } else {
        w.WriteHeader(http.StatusConflict)
    }
    return w
}

func getBookHandle(w http.ResponseWriter) http.ResponseWriter {
    books := AllBooks()
    writeJSON(w, books)
    return w
}

// BookHandleFunc to be used as http.HandleFunc for Book API
func BookHandleFunc(w http.ResponseWriter, r *http.Request) {
    isbn := r.URL.Path[len("/api/books/"):]

    switch method := r.Method; method {
    case http.MethodGet:
        book, found := GetBook(isbn)
        if found {
            w.WriteHeader(http.StatusOK)
            writeJSON(w, book)
        } else {
            w.WriteHeader(http.StatusNotFound)
        }
        
    case http.MethodPut:
        body, err := ioutil.ReadAll(r.Body)
        if hasInternalServerError(err, w) {
            return
        }
        book := FromJSON(body)
        exists := UpdateBook(isbn, book)
        if exists {
            w.WriteHeader(http.StatusOK)
        } else {
            w.WriteHeader(http.StatusNotFound)
        }
        
    case http.MethodDelete:
        DeleteBook(isbn)
        w.WriteHeader(http.StatusOK)
    
    default:
        w.WriteHeader(http.StatusBadRequest)
        w.Write([]byte("Unsupported request method."))
    }
}

func hasInternalServerError(err error, w http.ResponseWriter) bool {
    if err != nil {
        w.WriteHeader(http.StatusInternalServerError)
        return true
    }
    return false
}

func writeJSON(w http.ResponseWriter, i interface{}) {
    b, err := json.Marshal(i)
    if err != nil {
        panic(err)
    }
    w.Header().Add("Content-Type", "application/json; charset=utf-8")
    w.Write(b)
}


