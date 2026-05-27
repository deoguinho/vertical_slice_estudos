package list_users

import (
	"encoding/json"
	"net/http"

	"go-vertical-slice/internal/users/shared"
)

func Handler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(shared.Users)
}
