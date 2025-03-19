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

	r.Use(func(c *gin.Context) {
        c.Writer.Header().Set("Access-Control-Allow-Origin", cfg.Front.URL)
        c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
        c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Accept, Authorization")
        c.Writer.Header().Set("Access-Control-Expose-Headers", "Content-Disposition")

        c.Next()
    })

	r.POST("/postSQL", handler.PostSQLHandler)

	r.POST("/postCSV", handler.PostCSVHandler)

	log.Printf("Server running on port " + cfg.Server.Port)
	r.Run(":" + cfg.Server.Port)
}
