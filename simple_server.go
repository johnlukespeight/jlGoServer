package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

// User model
type User struct {
	gorm.Model
	Name     string    `json:"name"`
	Email    string    `json:"email"`
	Address  string    `json:"address"`
	JoinedAt time.Time `json:"joined_at"`
}

func main() {
	// Connect to database
	var err error
	db, err = gorm.Open(sqlite.Open("simple.db"), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	
	// Migrate schema
	db.AutoMigrate(&User{})
	
	// Set up routes
	http.HandleFunc("/users", usersHandler)
	
	// Start server
	fmt.Println("Server running on http://localhost:8888")
	log.Fatal(http.ListenAndServe(":8888", nil))
}

func usersHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		getUsers(w, r)
	case "POST":
		createUser(w, r)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func getUsers(w http.ResponseWriter, r *http.Request) {
	var users []User
	db.Find(&users)
	
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

func createUser(w http.ResponseWriter, r *http.Request) {
	var user User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	
	user.JoinedAt = time.Now()
	db.Create(&user)
	
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}
