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

	fmt.Printf("Server starting (%s)\n", port)
	fmt.Println("Routes:")
	for i := range routes {
		fmt.Printf("\t %s \n", routes[i])
	}

	fmt.Println("Ready")
	log.Fatal(http.ListenAndServe(port, nil))
}
