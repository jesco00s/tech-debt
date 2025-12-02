package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/jesco00s/tech-debt/resources"
)

func main() {
	const resourcesPath = "/resources/get"
	var routes = []string{resourcesPath}
	port := ":8080"

	http.HandleFunc(resourcesPath, resources.ResourcesHandler)

	for i := range routes {
		fmt.Println(routes[i])
	}

	fmt.Println("Server is running on port", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
