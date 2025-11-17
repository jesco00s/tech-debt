package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/jesco00s/tech-debt/items"
)

func itemsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		items.AddItem(w, r)
	} else {
		http.Error(w, "Unsupported method", http.StatusMethodNotAllowed)
	}
}

func main() {
	port := ":8080"
	http.HandleFunc("/items", itemsHandler)
	fmt.Println("Server is running on port", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
