package main

import (
	"fmt"
	"log"
	"github.com/abelmalu/CafeteriaAccessControl/internal/app"
	"github.com/joho/godotenv"
)

func main() {

    fmt.Println("before loading .env files")
	errEnv := godotenv.Load()
	fmt.Println("here")

	if errEnv != nil {

		fmt.Println(errEnv)
	}
	// Initialize the application: loads config, connects DB, sets up routes
	application, err := app.NewApp()
	if err != nil {
		log.Fatalf("Application startup failed: %v", err)
	}

	// Start the HTTP server
	application.Run()

}
