package controllers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/varconstant/go-bookstore/pkg/models"
	"github.com/varconstant/go-bookstore/pkg/utils"
	"net/http"
	"strconv"
)

var NewBook models.Book

func GetBook(w http.ResponseWriter, r *http.Request) {
	newBooks := models.GetAllBooks()
	res, _ := json.Marshal(newBooks)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(res)

}

func GetBooksById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	id, err := strconv.ParseInt(bookId, 10, 0)
	if err != nil {
		return
	}
	book, _ := models.GetBookById(id)
	res, _ := json.Marshal(book)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(res)
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	createBook := new(models.Book)
	utils.ParseBody(r, createBook)
	b := createBook.CreateBook()
	res, _ := json.Marshal(b)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(res)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	id, err := strconv.ParseInt(bookId, 10, 0)
	if err != nil {
		return
	}
	book := models.DeleteBookById(id)
	res, _ := json.Marshal(&book)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	_, err = w.Write(res)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	updateBook := new(models.Book)
	utils.ParseBody(r, updateBook)
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	id, err := strconv.ParseInt(bookId, 10, 0)
	if err != nil {
		return
	}
	bookDetails, db := models.GetBookById(id)
	if updateBook.Name != "" {
		bookDetails.Name = updateBook.Name
	}
	if updateBook.Author != "" {
		bookDetails.Author = updateBook.Author
	}
	if updateBook.Publication != "" {
		bookDetails.Publication = updateBook.Publication
	}
	db.Save(bookDetails)

	res, _ := json.Marshal(bookDetails)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	_, err = w.Write(res)
}
