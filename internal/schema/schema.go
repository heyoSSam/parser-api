package schema

import (
	"fmt"
	"parser-api/internal/reader"
	"regexp"
	"strconv"
	"strings"
)

func Inserts() {
	var text, err = reader.ReadPDF("../1.pdf")
	if err != nil {
		panic(err)
	}
	var sqlStatements []string

	re := regexp.MustCompile(`\s+`)
	text = re.ReplaceAllString(text, " ")

	sentences := regexp.MustCompile(`[.!?]\s+`).Split(text, -1)

	codeRegex := regexp.MustCompile(`(?s)^.*?КОДЕКС РЕСПУБЛИКИ КАЗАХСТАН`)
	partRegex := regexp.MustCompile(`ЧАСТЬ\s+(\d+)\.\s+(.+)`)
	sectionRegex := regexp.MustCompile(`РАЗДЕЛ\s+(\d+)\.\s+(.+)`)
	paragraphRegex := regexp.MustCompile(`ПАРАГРАФ\s+(\d+)\.\s+(.+)`)
	articleRegex := regexp.MustCompile(`Статья\s+(\d+)\.\s+(.+)`)
	clauseRegex := regexp.MustCompile(`(?m)^(\d+)\.\s+(.+)$`)
	subClauseRegex := regexp.MustCompile(`(?m)^(\d+)\)\s+(.+)$`)

	var codeID, partID, sectionID, paragraphID, articleID, clauseID, subClauseID int

	for _, line := range sentences {
		line = strings.TrimSpace(line)

		if matches := codeRegex.FindStringSubmatch(line); len(matches) > 0 {
			codeTitle := matches[0]
			sql := fmt.Sprintf(`INSERT INTO codes (code_id, code, version, status, actual_date, previous_date) VALUES (1, %s, 1, TRUE, NOW(), NULL) RETURNING code_id;`, codeTitle)
			sqlStatements = append(sqlStatements, sql)
			codeID = 1
		} else if matches := partRegex.FindStringSubmatch(line); len(matches) > 0 {
			partNum, partTitle := matches[1], matches[2]
			sql := fmt.Sprintf(`INSERT INTO part (part_id, code_id, code, version, status, actual_date, previous_date) VALUES (%d, %d, %s, 1, TRUE, NOW(), NULL) RETURNING part_id;`, partNum, codeID, partTitle)
			sqlStatements = append(sqlStatements, sql)
			partID, err = strconv.Atoi(partNum)
		} else if matches := sectionRegex.FindStringSubmatch(line); len(matches) > 0 {
			fmt.Println(matches)
			sectionNum, sectionTitle := matches[1], matches[2]
			sql := fmt.Sprintf(`INSERT INTO sections (section_number, code_id, part_id, section, version, status, actual_date, previous_date) VALUES (%d, %d, %d, %s, 1, TRUE, NOW(), NULL) RETURNING section_number;`, sectionNum, codeID, partID, sectionTitle)
			sqlStatements = append(sqlStatements, sql)
			sectionID, err = strconv.Atoi(sectionNum)
		} else if matches := paragraphRegex.FindStringSubmatch(line); len(matches) > 0 {
			paragraphNum, paragraphTitle := matches[1], matches[2]
			sql := fmt.Sprintf(`INSERT INTO paragraphs (paragraph_id, code_id, part_id, section_number, paragraph, version, status, actual_date, previous_date) VALUES (%d, %d, %d, %d, %s, 1, TRUE, NOW(), NULL) RETURNING paragraph_id;`, paragraphNum, codeID, partID, sectionID, paragraphTitle)
			sqlStatements = append(sqlStatements, sql)
			paragraphID, err = strconv.Atoi(paragraphNum)
		} else if matches := articleRegex.FindStringSubmatch(line); len(matches) > 0 {
			articleNum, articleTitle := matches[1], matches[2]
			sql := fmt.Sprintf(`INSERT INTO articles (article_number, code_id, part_id, section_number, paragraph_id, article, version, status, actual_date, previous_date) VALUES (%d, %d, %d, %d, %d, %s, 1, TRUE, NOW(), NULL) RETURNING article_number;`, articleNum, codeID, partID, sectionID, paragraphID, articleTitle)
			sqlStatements = append(sqlStatements, sql)
			articleID, err = strconv.Atoi(articleNum)
		} else if matches := clauseRegex.FindStringSubmatch(line); len(matches) > 0 {
			clauseNum, clauseTitle := matches[1], matches[2]
			sql := fmt.Sprintf(`INSERT INTO clauses (clause_number, code_id, part_id, section_number, paragraph_id, article_number, clause, version, status, actual_date, previous_date) VALUES (%d, %d, %d, %d, %d, %d, %s, 1, TRUE, NOW(), NULL) RETURNING clause_number;`, clauseNum, codeID, partID, sectionID, paragraphID, articleID, clauseTitle)
			sqlStatements = append(sqlStatements, sql)
			clauseID, err = strconv.Atoi(clauseNum)
		} else if matches := subClauseRegex.FindStringSubmatch(line); len(matches) > 0 {
			subClauseNum, subClauseTitle := matches[1], matches[2]
			subClauseID, err = strconv.Atoi(subClauseNum)
			sql := fmt.Sprintf(`INSERT INTO sub_clauses (sub_clause_number, code_id, part_id, section_number, paragraph_id, article_number, clause_number, sub_clause, version, status, actual_date, previous_date) VALUES (%d, %d, %d, %d, %d, %d, %d, %s, 1, TRUE, NOW(), NULL) RETURNING subClauseID;`, subClauseID, codeID, partID, sectionID, paragraphID, articleID, clauseID, subClauseTitle)
			sqlStatements = append(sqlStatements, sql)
		}
	}

	for _, sql := range sqlStatements {
		fmt.Println(sql)
	}
}
