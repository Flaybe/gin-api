package main

import (
	"log"

	"gin-api/config"
	"gin-api/internal/api"
)

func main() {
	log.Printf("Starting server in %s mode on %s\n", config.Env, config.ServerPort)

	r := api.NewRouter()
	if err := r.Run(config.ServerPort); err != nil {
		log.Fatalf("server exited: %v", err)
	}
}

