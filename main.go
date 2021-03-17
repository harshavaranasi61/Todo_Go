package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"gorm.io/gorm"
)

type Todo struct {
	gorm.Model
	Title  string `json:"Title"`
	Status string `json:"Status"`
}

func homepage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome to TODO app")
	fmt.Println("Homepage endpoint hit")
}

func returnAllTodos(w http.ResponseWriter, r *http.Request) {
	db := dbConfig()
	fmt.Println("Endpoint Hit: returnAllTodos")

	var todosList []Todo

	db.Find(&todosList)
	json.NewEncoder(w).Encode(todosList)

}

func handleRequests() {
	router := Router()
	//http.HandleFunc("/todo/add", addTodo)
	fmt.Println("Listening on port 8081")
	log.Fatal(http.ListenAndServe(":8081", router))

}

func main() {
	db := dbConfig()

	firstTodo := Todo{Title: "Sleep", Status: "false"}
	result := db.Create(&firstTodo)
	fmt.Println(firstTodo.ID, result)
	handleRequests()
}
