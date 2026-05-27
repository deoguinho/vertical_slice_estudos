package main

import (
	"log"
	"net/http"

	"go-vertical-slice/internal/users/create_user"
	"go-vertical-slice/internal/users/delete_user"
	"go-vertical-slice/internal/users/list_users"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("POST /users", create_user.Handler)
	mux.HandleFunc("GET /users/list", list_users.Handler)
	mux.HandleFunc("DELETE /users/{id}", delete_user.Handler)

	log.Println("Server running on: 8080")
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatal(err)
	}

}
