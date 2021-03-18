package main

import (
	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/", homepage).Methods("GET")
	router.HandleFunc("/todos", returnAllTodos).Methods("GET")
	router.HandleFunc("/todos/add", addNewTodo).Methods("POST")
	router.HandleFunc("/todos/complete/{id}", completeTodo).Methods("PUT")
	router.HandleFunc("/todos/undo/{id}", undoCompleteTodo).Methods("PUT")
	router.HandleFunc("/todos/remove/{id}", removeTodo).Methods("DELETE")
	router.HandleFunc("/todos/removeAll", removeAllTodos).Methods("DELETE")

	return router
}
