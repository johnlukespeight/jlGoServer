package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"../db"
	"../models"
)

// GetUsers returns all users
func GetUsers(w http.ResponseWriter, r *http.Request) {
	var users []models.User

	result := db.GetDB().Find(&users)
	if result.Error != nil {
		http.Error(w, "Error fetching users", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

// GetUser returns a specific user
func GetUser(w http.ResponseWriter, r *http.Request) {
	// Extract ID from the URL path
	paths := strings.Split(r.URL.Path, "/")
	if len(paths) < 3 {
		http.Error(w, "User ID not provided", http.StatusBadRequest)
		return
	}

	id, err := strconv.ParseUint(paths[len(paths)-1], 10, 64)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	var user models.User
	result := db.GetDB().First(&user, id)
	if result.Error != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

// CreateUser creates a new user
func CreateUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	result := db.GetDB().Create(&user)
	if result.Error != nil {
		http.Error(w, "Error creating user", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}

// UpdateUser updates an existing user
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	// Extract ID from the URL path
	paths := strings.Split(r.URL.Path, "/")
	if len(paths) < 3 {
		http.Error(w, "User ID not provided", http.StatusBadRequest)
		return
	}

	id, err := strconv.ParseUint(paths[len(paths)-1], 10, 64)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	var user models.User
	result := db.GetDB().First(&user, id)
	if result.Error != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	var updatedUser models.User
	err = json.NewDecoder(r.Body).Decode(&updatedUser)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Update fields
	if updatedUser.Name != "" {
		user.Name = updatedUser.Name
	}
	if updatedUser.Email != "" {
		user.Email = updatedUser.Email
	}
	if updatedUser.Address != "" {
		user.Address = updatedUser.Address
	}

	db.GetDB().Save(&user)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

// DeleteUser deletes a user
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	// Extract ID from the URL path
	paths := strings.Split(r.URL.Path, "/")
	if len(paths) < 3 {
		http.Error(w, "User ID not provided", http.StatusBadRequest)
		return
	}

	id, err := strconv.ParseUint(paths[len(paths)-1], 10, 64)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	var user models.User
	result := db.GetDB().First(&user, id)
	if result.Error != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	db.GetDB().Delete(&user)

	w.WriteHeader(http.StatusNoContent)
}
