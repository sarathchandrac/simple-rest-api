package main

// import (
// 	"fmt"
// 	"log"
// 	"net/http"
// 	"math/rand"
// 	"strconv"
// )
import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Book Struct (Model)
type Book struct {
	ID string `json:"id" json`
	ISBN string `json:"isbn" json`
	Title string `json:"title" json`
	Author *Author `json:"author" json`
}

// Author struct
type Author struct{
	FirstName string `json:"firstName" json`
	LastName string `json:"lastName" json`
}

// Get All Books
func getBooks(w http.ResponseWriter, r *http.Request){
}

// Get Book
func getBook(w http.ResponseWriter, r *http.Request)  {
}

// Create Book
func createBook(w http.ResponseWriter, r *http.Request){}

// Update Book
func updateBook(w http.ResponseWriter, r *http.Request){}

// Delete Book
func deleteBook(w http.ResponseWriter, r *http.Request){}

func main() {
	fmt.Println("Hello World!")
	r := mux.NewRouter()
	// Endpoints
	r.HandleFunc("/api/books", getBooks).Methods("GET")
	r.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	r.HandleFunc("/api/books", createBook).Methods("POST")
	r.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
	r.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8000", r))
}