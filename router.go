package main

import (
	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/", homepage).Methods("GET")
	router.HandleFunc("/todos", returnAllTodos).Methods("GET")
	router.HandleFunc("/todos/add", addNewTodo).Methods("POST")

	return router
}
