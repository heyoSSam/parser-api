package csv

import (
	"encoding/csv"
	"fmt"
	"os"
	"parser-api/internal/processing"
	"strings"

	"github.com/dlclark/regexp2"
)

type DocumentElement struct {
	Type  string
	ID    string
	Title string
}

func CreateCSVDump(docno string, text string, outputPath string) error {
	re := regexp2.MustCompile(`\s+`, 0)
	text, _ = re.Replace(text, " ", -1, -1)

	sentences := processing.SplitSentences(text)

	file, err := os.Create(outputPath)
	if err != nil {
		return fmt.Errorf("не удалось создать CSV файл: %w", err)
	}
	defer file.Close()

	file.WriteString("\xEF\xBB\xBF")

	writer := csv.NewWriter(file)
	writer.Comma = ';'
	defer writer.Flush()

	if err := writer.Write([]string{"Тип", "Номер", "Название"}); err != nil {
		return fmt.Errorf("ошибка записи заголовков: %w", err)
	}

	codeTitle := processing.IdToTitle[docno]
	if err := writer.Write([]string{"Кодекс", docno, codeTitle}); err != nil {
		return fmt.Errorf("ошибка записи названия кодекса: %w", err)
	}

	elements := parseDocumentElements(sentences)

	for _, elem := range elements {
		if err := writer.Write([]string{elem.Type, elem.ID, elem.Title}); err != nil {
			return fmt.Errorf("ошибка записи данных: %w", err)
		}
	}

	return nil
}

func parseDocumentElements(sentences []string) []DocumentElement {
	var elements []DocumentElement

	for _, line := range sentences {
		line = strings.TrimSpace(line)

		if m, _ := processing.CodeRegex.FindStringMatch(line); m != nil {
			codeTitle := m.String()
			elements = append(elements, DocumentElement{
				Type:  "Кодекс",
				ID:    "1",
				Title: codeTitle,
			})
		} else if m, _ := processing.PartRegex.FindStringMatch(line); m != nil {
			partNum := m.GroupByNumber(1).String()
			partTitle := m.GroupByNumber(2).String()
			elements = append(elements, DocumentElement{
				Type:  "Часть",
				ID:    partNum,
				Title: partTitle,
			})
		} else if m, _ := processing.SectionRegex.FindStringMatch(line); m != nil {
			sectionNum := m.GroupByNumber(1).String()
			sectionTitle := m.GroupByNumber(2).String()
			elements = append(elements, DocumentElement{
				Type:  "Раздел",
				ID:    sectionNum,
				Title: sectionTitle,
			})
		} else if m, _ := processing.ChapterRegex.FindStringMatch(line); m != nil {
			chapterNum := m.GroupByNumber(1).String()
			chapterTitle := m.GroupByNumber(2).String()
			elements = append(elements, DocumentElement{
				Type:  "Глава",
				ID:    chapterNum,
				Title: chapterTitle,
			})
		} else if m, _ := processing.ParagraphRegex.FindStringMatch(line); m != nil {
			paragraphNum := m.GroupByNumber(1).String()
			paragraphTitle := m.GroupByNumber(2).String()
			elements = append(elements, DocumentElement{
				Type:  "Параграф",
				ID:    paragraphNum,
				Title: paragraphTitle,
			})
		} else if m, _ := processing.ArticleRegex.FindStringMatch(line); m != nil {
			articleNum := m.GroupByNumber(1).String()
			articleTitle := m.GroupByNumber(2).String()
			elements = append(elements, DocumentElement{
				Type:  "Статья",
				ID:    articleNum,
				Title: articleTitle,
			})
		} else if m, _ := processing.ClauseRegex.FindStringMatch(line); m != nil {
			clauseNum := m.GroupByNumber(1).String()
			clauseTitle := m.GroupByNumber(2).String()
			elements = append(elements, DocumentElement{
				Type:  "Пункт",
				ID:    clauseNum,
				Title: clauseTitle,
			})
		}
	}

	return elements
}
