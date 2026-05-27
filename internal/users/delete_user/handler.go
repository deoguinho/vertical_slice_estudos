package delete_user

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"go-vertical-slice/internal/users/shared"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(r.URL.Path, "/")
	if len(parts) < 3 {
		http.Error(w, "User ID is required", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(parts[2])
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	found := false
	newUsers := make([]shared.User, 0)
	for _, user := range shared.Users {
		if user.ID == id {
			found = true
			continue
		}
		newUsers = append(newUsers, user)
	}

	if !found {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}
	shared.Users = newUsers

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "User deleted successfully"})
}
