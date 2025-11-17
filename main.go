package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/items", itemsHandler)
	fmt.Println("Hello World")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func itemsHandler(w http.ResponseWriter, r *http.Request) {

}
