package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/jesco00s/tech-debt/items"
	"github.com/jesco00s/tech-debt/resources"
)

func itemsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		items.AddItem(w, r)
	} else {
		http.Error(w, "Unsupported method", http.StatusMethodNotAllowed)
	}
}

func resourcesHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		resources.GetResources(w, r)
	} else {
		http.Error(w, "Unsupported method", http.StatusMethodNotAllowed)
	}
}

func main() {
	const resources_path = "/resources/get"
	var routes = []string{resources_path}
	port := ":8080"

	http.HandleFunc("/items", itemsHandler)
	http.HandleFunc(resources_path, resourcesHandler)
	log.Fatal(http.ListenAndServe(port, nil))

	for i := range routes {
		fmt.Println(routes[i])
	}

	fmt.Println("Server is running on port", port)
}
