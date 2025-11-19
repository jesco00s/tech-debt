package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/jesco00s/tech-debt/resources"
)

func main() {
	const resources_path = "/resources/get"
	var routes = []string{resources_path}
	port := ":8080"

	http.HandleFunc(resources_path, resources.ResourcesHandler)

	for i := range routes {
		fmt.Println(routes[i])
	}

	fmt.Println("Server is running on port", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
