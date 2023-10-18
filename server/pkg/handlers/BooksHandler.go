package handlers

import (
	"bookstore-go/server/pkg/models"
	"bookstore-go/server/pkg/utils"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

var NewBook models.Book;

var CreateBookHandler = func(w http.ResponseWriter, r *http.Request) {
	newBook:= &models.Book{};
	utils.ParseBody(r, newBook);
	b:= newBook.CreateBook();
	res,_:= json.Marshal(b);
	w.WriteHeader(http.StatusOK);
	w.Write(res);	
}

var GetBooksHandler = func(w http.ResponseWriter, r *http.Request) {
	
	allBooks := models.GetAllBooks();
	res, _:= json.Marshal(allBooks);
	w.Header().Set("Content-Type", "pkglication/json");
	w.WriteHeader(http.StatusOK);
	w.Write(res);
}

var GetSingleBookHandler = func(w http.ResponseWriter, r *http.Request) {
	vars:= mux.Vars(r);
	bookId:= vars["bookId"];
	Id,err:= strconv.ParseInt(bookId,0,0);
	if err!=nil{
		http.Error(w, "Unable to parse Id", http.StatusBadRequest);
		return;
	}
	bookDetails, _ :=models.GetBookById(Id);
	res, _:= json.Marshal(bookDetails);
	w.Header().Set("Content-Type", "pkglication/json");
	w.WriteHeader(http.StatusOK);
	w.Write(res);
}

var UpdateBookHandler = func(w http.ResponseWriter, r *http.Request) {
	var updatedBook =  &models.Book{};
	utils.ParseBody(r, updatedBook);
	vars:=mux.Vars(r);
	bookId:=vars["bookId"];
	Id,err:= strconv.ParseInt(bookId,0,0);
	if err!=nil{
		http.Error(w, "Cannot parse Id", http.StatusBadRequest);
		return;
	}
	bookDetails, db:=models.GetBookById(Id);
	if updatedBook.Name != ""{
		bookDetails.Name = updatedBook.Name;
	}
	if updatedBook.Author != ""{
		bookDetails.Author = updatedBook.Author;
	}
	if updatedBook.Publication != ""{
		bookDetails.Publication = updatedBook.Publication;
	}
	db.Save(&bookDetails)
	res, _ := json.Marshal(bookDetails);
	w.Header().Set("Content-Type", "pkglication/json");
	w.WriteHeader(http.StatusOK);
	w.Write(res);
}

var DeleteBookHandler = func(w http.ResponseWriter, r *http.Request) {
	vars:= mux.Vars(r);
	bookId:=vars["bookId"];
	Id,err:= strconv.ParseInt(bookId,0,0);
	if err!=nil{
		http.Error(w, "Cannot parsed Id", http.StatusBadRequest);
		return;
	}
	book:= models.DeleteBook(Id);
	res,_:= json.Marshal(book);
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK);
	w.Write(res);
}