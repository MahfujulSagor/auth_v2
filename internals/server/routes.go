package server

import (
	"log"
	"net/http"
	"time"
)

// registerRoutes sets up the HTTP routes and middleware for the server.
func registerRoutes(mux *http.ServeMux) {
	// Health check endpoint
	mux.Handle("/health", logger(makeHandler(healthHandler)))

	// Additional routes
	mux.Handle("POST /users", logger(makeHandler(createUserHandler)))
	mux.Handle("PUT /users", logger(makeHandler(updateUserHandler)))
	mux.Handle("GET /users", logger(makeHandler(getUsersListHandler)))
	mux.Handle("GET /users/{id}", logger(makeHandler(getUserByIdHandler)))
	mux.Handle("DELETE /users/{id}", logger(makeHandler(deleteUserHandler)))
}

// ========== Types ==========
type apiHandler func(w http.ResponseWriter, r *http.Request)

// makeHandler converts a function with the signature of apiHandler to an http.HandlerFunc.
func makeHandler(fn apiHandler) http.HandlerFunc {
	return http.HandlerFunc(fn)
}

// ========== Handlers ==========

func createUserHandler(w http.ResponseWriter, r *http.Request) {
	// Placeholder for creating a user
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("User created"))
}

func updateUserHandler(w http.ResponseWriter, r *http.Request) {
	// Placeholder for updating a user
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("User updated"))
}

func getUsersListHandler(w http.ResponseWriter, r *http.Request) {
	// Placeholder for getting a list of users
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("List of users"))
}

func getUserByIdHandler(w http.ResponseWriter, r *http.Request) {
	// Placeholder for getting a user by ID
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("User details by ID"))
}

func deleteUserHandler(w http.ResponseWriter, r *http.Request) {
	// Placeholder for deleting a user
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("User deleted"))
}

// ========== Health Check ==========

// healthHandler responds with a simple "OK" message to indicate the server is running.
func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

// ========== Middleware ==========

// logger is a middleware that logs the details of each HTTP request.
func logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		flags := "| " + r.Method + " | " + r.URL.Path
		log.Println(flags, "|", time.Since(start))
	})
}

