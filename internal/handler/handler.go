package handler

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"parser-api/internal/processing"
	"parser-api/internal/reader"
	"parser-api/internal/schema"
)

type RequestBody struct {
	Docno string `json:"docno"`
}

func PostSQLHandler(c *gin.Context) {
	var req RequestBody
	if err := c.BindJSON(&req); err != nil {
		log.Println("Failed to parse JSON request:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	if req.Docno == "" {
		log.Println("Missing docno in request")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing docno"})
		return
	}

	sqlFilePath := "./output.sql"
	err := schema.WriteSQLToFile(sqlFilePath, schema.Inserts(req.Docno))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate SQL"})
		return
	}

	c.FileAttachment(sqlFilePath, "output.sql")
	defer os.Remove(sqlFilePath)
}

func PostCSVHandler(c *gin.Context) {
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

	text, err := reader.ReadPDF(filePath)
	if err != nil {
		log.Printf("Failed to read PDF: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read PDF"})
		return
	}

	csvFilePath := "./output.csv"
	err = processing.CreateMultiSheetCSV(text, csvFilePath)
	if err != nil {
		log.Printf("Failed to generate CSV: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate CSV"})
		return
	}

	c.FileAttachment(csvFilePath, "output.csv")

	defer func() {
		_ = os.Remove(filePath)
		_ = os.Remove(csvFilePath)
	}()
}
