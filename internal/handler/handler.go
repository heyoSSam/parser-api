package handler

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"parser-api/internal/schema"
)

func PostSQLHandler(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		log.Printf("Failed to get form file: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to get file"})
		return
	}

	filePath := "./" + file.Filename
	if err := c.SaveUploadedFile(file, filePath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save file"})
		return
	}

	sqlFilePath := "./output.sql"
	err = schema.WriteSQLToFile(sqlFilePath, schema.Inserts(filePath))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate SQL"})
		return
	}

	c.FileAttachment(sqlFilePath, "output.sql")

	defer func() {
		_ = os.Remove(filePath)
		_ = os.Remove(sqlFilePath)
	}()
}
