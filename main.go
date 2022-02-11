package main

import (
	"basic-api/handles"
	"log"
	"net/http"
)

func handleRequest() {
	log.Println("start Service")
	router := handles.Router()
	log.Fatal(http.ListenAndServe(":9002", router))
}
func main() {
	handleRequest()
}
