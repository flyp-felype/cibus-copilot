package main

import (
	"copilot/internal/config"
	"copilot/internal/router"
	"copilot/internal/server"
	"log"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("could not load config: %v", err)
	}

	r := router.New()
	srv := server.New(cfg, r)

	if err := srv.Run(); err != nil {
		log.Fatalf("server error: %v", err)
	}
}
