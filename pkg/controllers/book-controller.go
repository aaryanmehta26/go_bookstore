package controllers

import (
	"fmt"
	"net/http"
)

func CreateBook(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create Book")
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {

}

func GetBook(w http.ResponseWriter, r *http.Request) {}

func GetBookById(w http.ResponseWriter, r *http.Request) {}

func UpdateBook(w http.ResponseWriter, r *http.Request) {}
