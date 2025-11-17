package main

import (
	items "TechDebt/Items"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func itemsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		addItem(w, r)
	} else {
		http.Error(w, "Unsupported method", http.StatusMethodNotAllowed)
	}
}

func addItem(w http.ResponseWriter, r *http.Request) {
	var newItem items.Item
	err := json.NewDecoder(r.Body).Decode(&newItem)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}
	fmt.Println(newItem)
}

func main() {
	port := ":8080"
	http.HandleFunc("/items", itemsHandler)
	fmt.Println("Server is running on port", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
