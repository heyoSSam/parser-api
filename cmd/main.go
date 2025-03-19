package main

import (
	"parser-api/internal/schema"
)

func main() {
	err := schema.WriteSQLToFile("a", schema.Inserts("../1.pdf"))
	if err != nil {
		panic(err)
	}
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
