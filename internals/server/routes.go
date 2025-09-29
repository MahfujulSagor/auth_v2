package server

import (
	"log"
	"net/http"
	"time"
)

// registerRoutes sets up the HTTP routes and middleware for the server.
func registerRoutes(mux *http.ServeMux) {
	// Health check endpoint
	mux.Handle("/health", logger(http.HandlerFunc(healthHandler)))

	// Additional routes
}

// ========== Handlers ==========

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
		flags := " | " + r.Method + " | " + r.URL.Path
		log.Println(flags, "|", time.Since(start))
	})
}

