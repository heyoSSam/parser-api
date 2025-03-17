package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"parser-api/config"
	"parser-api/internal/reader"
)

func main() {
	cfg, err := config.LoadConfig("../config.yaml")
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	text, err := reader.ReadPDF("../")
	if err != nil {
		log.Fatalf("Failed to read pdf: %v", err)
	}
	fmt.Println(text)

	r := gin.Default()

	log.Printf("Server running on port " + cfg.Server.Port)
	r.Run(":" + cfg.Server.Port)
}
