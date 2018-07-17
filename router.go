package main

import (
	"github.com/gorilla/mux"
	d "github.com/user/go-test-rest/delivery"
)

// Define HTTP request routes
func newRouter() *mux.Router {
	r := mux.NewRouter()

	// user api
	r.HandleFunc("/user", d.AllUserEndPoint).Methods("GET")
	r.HandleFunc("/user", d.CreateUserEndPoint).Methods("POST")
	r.HandleFunc("/user", d.UpdateUserEndPoint).Methods("PUT")
	r.HandleFunc("/user", d.DeleteUserEndPoint).Methods("DELETE")
	r.HandleFunc("/user/{id}", d.FindUserEndpoint).Methods("GET")

	// user api
	r.HandleFunc("/book", d.GetAllBook).Methods("GET")
	r.HandleFunc("/book", d.CreateBookData).Methods("POST")
	r.HandleFunc("/book", d.UpdateBookData).Methods("PUT")
	r.HandleFunc("/book", d.DeleteBookData).Methods("DELETE")
	r.HandleFunc("/book/{id}", d.FindBookById).Methods("GET")

	return r
}
