package main

import (
	"log"
	"parser-api/config"
	"parser-api/internal/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg, err := config.LoadConfig("../config.yaml")
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	r := gin.Default()

	r.POST("/postSQL", handler.PostSQLHandler)

	r.POST("/postCSV", handler.PostCSVHandler)

	log.Printf("Server running on port " + cfg.Server.Port)
	r.Run(":" + cfg.Server.Port)
}
