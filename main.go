package main

import (
	"fmt"
	"log"
	"net/http"
)

func handleRequests() {
	router := Router()
	fmt.Println("Listening on port 8081")
	log.Fatal(http.ListenAndServe(":8081", router))
}

func main() {
	dbConfig()
	handleRequests()
}
