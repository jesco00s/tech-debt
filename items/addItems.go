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

	connStr := "user=postgres dbname=josejescobar host=localhost sslmode=disable"

	// Connect to the database
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	fmt.Println("err not equal to nil. can continue")

}
