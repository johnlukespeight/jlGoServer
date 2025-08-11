package main

import (
	"net/http"
	"strings"
)

// UserRouteHandler handles all user-related routes
func UserRouteHandler(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	
	// Route for listing all users or creating a new user
	if path == "/api/users" {
		switch r.Method {
		case http.MethodGet:
			GetUsers(w, r)
		case http.MethodPost:
			CreateUser(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
		return
	}
	
	// Routes for specific user operations (get, update, delete)
	if strings.HasPrefix(path, "/api/users/") {
		switch r.Method {
		case http.MethodGet:
			GetUser(w, r)
		case http.MethodPut:
			UpdateUser(w, r)
		case http.MethodDelete:
			DeleteUser(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
		return
	}
	
	// If no routes match
	http.NotFound(w, r)
}
