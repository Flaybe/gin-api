package main

import (
	"gin-api/config"
	"gin-api/internal/api"
	"log"
)

func main() {
	log.Printf("Starting server in %s mode on %s\n", config.Env, config.ServerPort)

	// Initialize and run the server
	router := api.NewRouter()
	if err := router.Run(config.ServerPort); err != nil {
		log.Fatalf("server exited: %v", err)
	}
}

