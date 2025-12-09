package main

import (
	"fmt"
	"log"

	"github.com/abelmalu/CafeteriaAccessControl/internal/app"

	"github.com/abelmalu/CafeteriaAccessControl/internal/redis"
	"github.com/joho/godotenv"
)

func main() {

	redisdb.InitRedis()

	log.Println("Starting server...")

	errEnv := godotenv.Load()

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
