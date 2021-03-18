package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

var Db = dbConfig()

func homepage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome to TODO app")
	fmt.Println("Homepage endpoint hit")
}

func returnAllTodos(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: returnAllTodos")

	var todosList []Todo

	Db.Find(&todosList)
	json.NewEncoder(w).Encode(todosList)

}

func addNewTodo(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: addNewTodo")

	var todo Todo
	_ = json.NewDecoder(r.Body).Decode(&todo)

	Db.Create(&todo)
	json.NewEncoder(w).Encode(todo)
}

func completeTodo(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: completeTodo")

	var todo Todo
	params := mux.Vars(r)
	Db.Where("id = ?", params["id"]).Find(&todo)
	todo.Status = "Completed"
	Db.Save(&todo)
	json.NewEncoder(w).Encode(todo)
}

func undoCompleteTodo(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: UndoCompleteTodo")

	var todo Todo
	params := mux.Vars(r)
	Db.Where("id = ?", params["id"]).Find(&todo)
	todo.Status = "False"
	Db.Save(&todo)
	json.NewEncoder(w).Encode(todo)
}

func removeTodo(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: removeTodo")

	var todo Todo
	params := mux.Vars(r)
	Db.Where("id = ?", params["id"]).Delete(&todo)
	json.NewEncoder(w).Encode(params["id"] + " is deleted")
}

func removeAllTodos(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: removeAllTodos")

	Db.Where("1 = 1").Delete(&Todo{})
	json.NewEncoder(w).Encode("All todos are removed")
}
