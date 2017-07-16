package main

import (
	"cloud-native-go/api"
	"fmt"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/echo", echo)
	http.HandleFunc("/books", api.BooksHandleFunc)
	http.ListenAndServe(port(), nil)
}

func port() string {
	port := os.Getenv("port")

	if len(port) == 0 {
		port = ":8080"
	}

	return port
}

func index(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Hello cloud native go")
}

func echo(w http.ResponseWriter, r *http.Request) {
	message := r.URL.Query()["message"][0]
	w.WriteHeader(http.StatusOK)
	w.Header().Add("Content-Type", "text/plain")
	fmt.Fprintf(w, message)
}