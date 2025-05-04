package controllers

import (
	"encoding/json"
	"fmt"
	"mysql/pkg/models"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func CreateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	var Books models.Book
	err := json.NewDecoder(r.Body).Decode(&Books)
	if err != nil {
		json.NewEncoder(w).Encode("Error in parsing json")
		return
	}
	b := Books.CreateBook()
	w.Header().Set("content-type", "application-json")
	json.NewEncoder(w).Encode(b)
}

func GetBook(w http.ResponseWriter, r *http.Request) {
	newBooks := models.GetALlBooks()
	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(newBooks)
}

func GetBookByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	bookDetails, _ := models.GetbookByID(ID)
	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(bookDetails)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookID := vars["bookId"]
	id, _ := strconv.ParseInt(bookID, 0, 0)
	b := models.DeleteBook(id)
	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(b)
}
