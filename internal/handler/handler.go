package handler

import (
	"log"
	"net/http"
	"os"
	"parser-api/internal/csv"
	"parser-api/internal/processing"
	"parser-api/internal/schema"

	"github.com/gin-gonic/gin"
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
	var req RequestBody
	if err := c.BindJSON(&req); err != nil {
		log.Println("Ошибка парсинга JSON запроса:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	if req.Docno == "" {
		log.Println("Отсутствует docno в запросе")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing docno"})
		return
	}

	text, err := processing.GetDocumentText(req.Docno)
	if err != nil {
		log.Println("Ошибка получения текста документа:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch document text"})
		return
	}

	csvFilePath := "./output.csv"
	err = csv.CreateCSVDump(text, csvFilePath)
	if err != nil {
		log.Printf("Ошибка генерации CSV: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate CSV"})
		return
	}

	c.FileAttachment(csvFilePath, "output.csv")
	defer os.Remove(csvFilePath)
}
