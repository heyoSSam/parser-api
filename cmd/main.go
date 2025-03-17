package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"parser-api/config"
)

func main() {
	cfg, err := config.LoadConfig("../config.yaml")
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	r := gin.Default()

	log.Printf("Server running on port " + cfg.Server.Port)
	r.Run(":" + cfg.Server.Port)
}
