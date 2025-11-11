package main

import (
	"github.com/abelmalu/CafeteriaAccessControl/internal/app"
	"log"
)

func main() {
	// Initialize the application: loads config, connects DB, sets up routes
	application, err := app.NewApp()
	if err != nil {
		log.Fatalf("Application startup failed: %v", err)
	}
	defer application.DB.Close()

	// Start the HTTP server
	application.Run()
}
