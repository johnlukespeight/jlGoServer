package main

import (
	"fmt"
	"log"
	"net/http"
)

func formHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}
	fmt.Fprintln(w, "POST request successful")
	name := r.FormValue("name")
	address := r.FormValue("address")
	fmt.Fprintf(w, "Name = %s\n", name)
	fmt.Fprintf(w, "Address = %s\n", address)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	if r.Method != "GET" {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}

	fmt.Fprintf(w, "Hello!")
}

func main() {
	// Initialize database connection
	InitDB()
	
	// Auto migrate models
	DB.AutoMigrate(&User{})

	// Create a new server mux for more control
	mux := http.NewServeMux()
	
	// Set up static file server
	fileServer := http.FileServer(http.Dir("./static"))
	mux.Handle("/", fileServer)
	
	// Register existing handlers
	mux.HandleFunc("/form", formHandler)
	mux.HandleFunc("/hello", helloHandler)
	
	// Register new API routes - make sure the more specific route comes first
	mux.HandleFunc("/api/users/", UserRouteHandler)  // This matches /api/users/{id}
	mux.HandleFunc("/api/users", UserRouteHandler)   // This matches exact /api/users
	
	// Wrap with debug handler
	handler := debugHandler(mux)

	fmt.Printf("Server started at port 8080\n")
	fmt.Printf("Access user management UI at http://localhost:8080/users.html\n")
	if err := http.ListenAndServe(":8080", handler); err != nil {
		log.Fatal(err)
	}
}
