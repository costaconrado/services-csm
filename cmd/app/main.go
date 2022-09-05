package main

import (
	"log"

	"github.com/costaconrado/services-csm/config"
	"github.com/costaconrado/services-csm/internal/app"
)

func main() {
	// Configuration
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}

	// Run
	app.Run(cfg)
}
