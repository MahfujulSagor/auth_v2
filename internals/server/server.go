package server

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/MahfujulSagor/auth_v2/internals/database"
	_ "github.com/joho/godotenv/autoload"
)

type Server struct {
	port int

	db database.Service
}

// NewServer initializes the server with configurations and returns an http.Server instance
func NewServer() *http.Server {
	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		fmt.Println("PORT not set in environment variables, defaulting to 8080")
		port = 8080
	}

	NewServer := &Server{
		port: port,
		db:   database.New(),
	}

	// Declare the HTTP server with timeouts
	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", NewServer.port),
		Handler:      NewServer.RegisterRoutes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	return server
}
