package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/aaryanmehta26/go_bookstore/pkg/models"
)

var NewBook models.Book

func CreateBook(w http.ResponseWriter, r *http.Request) {

}

func DeleteBook(w http.ResponseWriter, r *http.Request) {

}

func GetBook(w http.ResponseWriter, r *http.Request) {
	newBooks := models.GetAllBooks()
	res, _ := json.Marshal(newBooks) // convert to json
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetBookById(w http.ResponseWriter, r *http.Request) {}

func UpdateBook(w http.ResponseWriter, r *http.Request) {}
