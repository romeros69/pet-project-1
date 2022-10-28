package main

import (
	"log"
	"pet-project-1/config"
	"pet-project-1/internal/app"
)

// @tittle Simple Book CRUDL
// @version 1.0
// @description API Server for Book Application

// @host localhost:9000
// @BasePath /
func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Error in parse config: %s\n", err)
	}

	app.Run(cfg)
}
