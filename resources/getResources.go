package resources

import (
	"encoding/json"
	"fmt"

	"net/http"

	"github.com/jesco00s/tech-debt/constants"
	"github.com/jesco00s/tech-debt/db"
	_ "github.com/lib/pq"
)

func ResourcesHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		getResources(w)
	} else {
		http.Error(w, "Unsupported method", http.StatusMethodNotAllowed)
	}
}

func getResources(w http.ResponseWriter) {
	dbConn := db.GetDbConnection()
	defer dbConn.Close()

	selectResourceQuery := fmt.Sprintf("SELECT ar.resource_id, ar.resource_name, ar.os , ar.application_id, ar.product_id FROM %s ar", constants.All_Resources_Table_Name)
	rows, err := dbConn.Query(selectResourceQuery)
	if err != nil {
		http.Error(w, "Error getting resources", http.StatusInternalServerError)
		return
	}

	defer rows.Close()

	var resources []Resource

	for rows.Next() {
		var res Resource
		if err := rows.Scan(&res.ResourceId, &res.ResourceName, &res.Os, &res.ApplicationId, &res.ProductId); err != nil {
			http.Error(w, "Error scanning resource", http.StatusInternalServerError)
			return
		}
		resources = append(resources, res)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resources)
}
