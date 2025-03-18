package main

import (
	"parser-api/internal/schema"
)

func main() {
	schema.Inserts()
	//cfg, err := config.LoadConfig("../config.yaml")
	//if err != nil {
	//	log.Fatalf("Failed to load config: %v", err)
	//}
	//
	//r := gin.Default()
	//
	//log.Printf("Server running on port " + cfg.Server.Port)
	//r.Run(":" + cfg.Server.Port)
}
