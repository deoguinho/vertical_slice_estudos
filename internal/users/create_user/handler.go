package create_user

import (
	"encoding/json"
	"net/http"

	"go-vertical-slice/internal/users/shared"
)

func Handler(w http.ResponseWriter, r *http.Request) {

	var command CreateUserCommand

	err := json.NewDecoder(r.Body).Decode(&command)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	user := shared.User{
		ID:   len(shared.Users) + 1,
		Name: command.Name,
	}
	shared.Users = append(shared.Users, user)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "User created successfully",
	})
}
