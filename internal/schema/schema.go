package schema

import (
	"fmt"
	"github.com/dlclark/regexp2"
	"os"
	"parser-api/internal/processing"
	"parser-api/internal/reader"
	"strconv"
	"strings"
)

func Inserts(path string) []string {
	text, err := reader.ReadPDF(path)
	if err != nil {
		panic(err)
	}
	var sqlStatements []string

	re := regexp2.MustCompile(`\s+`, 0)
	text, _ = re.Replace(text, " ", -1, -1)

	sentences := processing.SplitSentences(text)

	var codeID, partID, sectionID, paragraphID, articleID, clauseID int
	clauseCount := 0

	for _, line := range sentences {
		line = strings.TrimSpace(line)

		if m, _ := processing.CodeRegex.FindStringMatch(line); m != nil {
			codeTitle := m.String()
			sql := fmt.Sprintf(`UPDATE codes SET code = '%s', actual_date = NOW() WHERE code_id = 1;`, codeTitle)
			sqlStatements = append(sqlStatements, sql)
			codeID = 1
		} else if m, _ := processing.PartRegex.FindStringMatch(line); m != nil {
			partNum, _ := strconv.Atoi(m.GroupByNumber(1).String())
			partTitle := m.GroupByNumber(2).String()
			partID = partNum
			sql := fmt.Sprintf(`INSERT INTO part (part_id, code_id, code, version, status, actual_date, previous_date) VALUES (%d, %d, '%s', 1, TRUE, NOW(), NULL);`, partID, codeID, partTitle)
			sqlStatements = append(sqlStatements, sql)
		} else if m, _ := processing.SectionRegex.FindStringMatch(line); m != nil {
			sectionNum, _ := strconv.Atoi(m.GroupByNumber(1).String())
			sectionTitle := m.GroupByNumber(2).String()
			sectionID = sectionNum
			sql := fmt.Sprintf(`INSERT INTO sections (section_number, code_id, part_id, section, version, status, actual_date, previous_date) VALUES (%d, %d, %d, '%s', 1, TRUE, NOW(), NULL);`, sectionID, codeID, partID, sectionTitle)
			sqlStatements = append(sqlStatements, sql)
		} else if m, _ := processing.ParagraphRegex.FindStringMatch(line); m != nil {
			paragraphNum, _ := strconv.Atoi(m.GroupByNumber(1).String())
			paragraphTitle := m.GroupByNumber(2).String()
			paragraphID = paragraphNum
			sql := fmt.Sprintf(`INSERT INTO paragraphs (paragraph_id, code_id, part_id, section_number, paragraph, version, status, actual_date, previous_date) VALUES (%d, %d, %d, %d, '%s', 1, TRUE, NOW(), NULL);`, paragraphID, codeID, partID, sectionID, paragraphTitle)
			sqlStatements = append(sqlStatements, sql)
		} else if m, _ := processing.ArticleRegex.FindStringMatch(line); m != nil {
			articleNum, _ := strconv.Atoi(m.GroupByNumber(1).String())
			articleTitle := m.GroupByNumber(2).String()
			articleID = articleNum
			sql := fmt.Sprintf(`INSERT INTO articles (article_number, code_id, part_id, section_number, paragraph_id, article, version, status, actual_date, previous_date) VALUES (%d, %d, %d, %d, %d, '%s', 1, TRUE, NOW(), NULL);`, articleID, codeID, partID, sectionID, paragraphID, articleTitle)
			sqlStatements = append(sqlStatements, sql)
		} else if m, _ := processing.ClauseRegex.FindStringMatch(line); m != nil {
			clauseNum, _ := strconv.Atoi(m.GroupByNumber(1).String())
			clauseTitle := m.GroupByNumber(2).String()
			clauseID = clauseNum
			clauseCount++;
			sql := fmt.Sprintf(`INSERT INTO clauses (clause_id, clause_number, code_id, part_id, section_number, paragraph_id, article_number, clause, version, status, actual_date, previous_date) VALUES (%d, %d, %d, %d, %d, %d, %d, '%s', 1, TRUE, NOW(), NULL);`, clauseCount, clauseID, codeID, partID, sectionID, paragraphID, articleID, clauseTitle)
			sqlStatements = append(sqlStatements, sql)
		}
	}

	return sqlStatements
}

func WriteSQLToFile(filename string, sqlStatements []string) error {
	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}
	defer file.Close()

	_, err = file.WriteString(Skeleton + "\n\n")
	if err != nil {
		return fmt.Errorf("failed to write schema to file: %w", err)
	}

	for _, sql := range sqlStatements {
		_, err := file.WriteString(sql + "\n")
		if err != nil {
			return fmt.Errorf("failed to write inserts to file: %w", err)
		}
	}

	return nil
}
