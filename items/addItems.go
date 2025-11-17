package items

import (
	"database/sql"
	"encoding/json"

	"fmt"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

func AddItem(w http.ResponseWriter, r *http.Request) {
	var newItem Item
	err := json.NewDecoder(r.Body).Decode(&newItem)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	//todo need to separate this into its own function. Do not like this but it is ok for now
	connStr := "user=postgres dbname=josejescobar host=localhost sslmode=disable"

	// Connect to the database
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	fmt.Println("err not equal to nil. can continue")

	// Insert the item into the database
	_, err = db.Exec("INSERT INTO public.tech_debt (resource_name, tech_debt_id) VALUES ($1, $2)", newItem.ResourceName, newItem.TechDebtId)
	if err != nil {
		http.Error(w, "Error saving item", http.StatusInternalServerError)
		return
	}

}
