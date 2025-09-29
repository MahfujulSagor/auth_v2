package main

import "github.com/MahfujulSagor/auth_v2/internals/server"

func main() {
	// Initialize and start the server
	server := server.New(":8080")
	server.Start()
}
