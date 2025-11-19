package resources

import (
	"database/sql"
	"encoding/json"

	"log"
	"net/http"

	_ "github.com/lib/pq"
)

func ResourcesHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		GetResources(w)
	} else {
		http.Error(w, "Unsupported method", http.StatusMethodNotAllowed)
	}
}

func GetResources(w http.ResponseWriter) {

	//todo need to separate this into its own function. Do not like this but it is ok for now
	connStr := "user=postgres dbname=josejescobar host=localhost sslmode=disable"

	// Connect to the database
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Insert the item into the database
	const selectResourceQuery = "Select ar.resource_id, ar.resource_name, ar.os , ar.application_id, ar.product_id from public.all_resources ar "
	rows, err := db.Query(selectResourceQuery)
	if err != nil {
		http.Error(w, "Error getting resources", http.StatusInternalServerError)
		return
	}

	defer rows.Close()

	var resources []Resource

	// Iterate over the rows and scan into the struct
	for rows.Next() {
		var res Resource
		if err := rows.Scan(&res.ResourceId, &res.ResourceName, &res.Os, &res.ApplicationId, &res.ProductId); err != nil {
			http.Error(w, "Error scanning resource", http.StatusInternalServerError)
			return
		}
		resources = append(resources, res)
	}

	// Check for errors from iterating over rows
	if err := rows.Err(); err != nil {
		http.Error(w, "Error during row iteration", http.StatusInternalServerError)
		return
	}

	// Set the response header to application/json
	w.Header().Set("Content-Type", "application/json")
	// Respond with the resources in JSON format
	json.NewEncoder(w).Encode(resources)

}
