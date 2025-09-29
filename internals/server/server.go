package server

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type Server struct {
	httpServer *http.Server
}

func New(addr string) *Server {
	mux := http.NewServeMux()
	// Register routes
	registerRoutes(mux) // Defined in routes.go

	server := &http.Server{
		Addr:         addr,
		Handler:      mux,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	return &Server{
		httpServer: server,
	}
}

// Start runs the HTTP server and handles graceful shutdown on interrupt signals.
func (s *Server) Start() {
	// Start the server in a goroutine
	go func() {
		log.Println("Starting server on", s.httpServer.Addr)
		if err := s.httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal("Server failed: ", err)
		}
	}()

	// Create a channel to listen for interrupt or terminate signals from the OS.
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)
	<-stop

	// Create a context to attempt a graceful 5 second shutdown.
	ctx, calcel := context.WithTimeout(context.Background(), 5*time.Second)
	defer calcel()

	log.Println("Shutting down server...")
	if err := s.httpServer.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown: ", err)
	}
	log.Println("Server stopped gracefully")
}
