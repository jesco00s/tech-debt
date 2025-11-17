package items

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func AddItem(w http.ResponseWriter, r *http.Request) {
	var newItem items.Item
	err := json.NewDecoder(r.Body).Decode(&newItem)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}
	fmt.Println(newItem)
}
