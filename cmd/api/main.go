package main

import (
	"log"

	"github.com/abelmalu/CafeteriaAccessControl/internal/app"
)

func main() {

	// errEnv := godotenv.Load()

	// if errEnv != nil {

	// 	fmt.Println(errEnv)
	// }
	// Initialize the application: loads config, connects DB, sets up routes
	application, err := app.NewApp()
	if err != nil {
		log.Fatalf("Application startup failed: %v", err)
	}

	// Start the HTTP server
	application.Run()

}
