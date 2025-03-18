package schema

import (
	"fmt"
	"parser-api/internal/processing"
	"parser-api/internal/reader"
	"regexp"
	"strconv"
	"strings"
)

func Inserts() {
	var text, err = reader.ReadPDF("..\\1.pdf")
	if err != nil {
		panic(err)
	}
	var sqlStatements []string

	re := regexp.MustCompile(`\s+`)
	text = re.ReplaceAllString(text, " ")

	sentences := processing.SplitSentences(text)

	var codeID, partID, sectionID, paragraphID, articleID, clauseID, subClauseID int
	for _, line := range sentences {
		line = strings.TrimSpace(line)

		if matches := processing.CodeRegex.FindStringSubmatch(line); len(matches) > 0 {
			codeTitle := matches[0]
			sql := fmt.Sprintf(`INSERT INTO codes (code_id, code, version, status, actual_date, previous_date) VALUES (1, %s, 1, TRUE, NOW(), NULL) RETURNING code_id;`, codeTitle)
			sqlStatements = append(sqlStatements, sql)
			codeID = 1
		} else if matches := processing.PartRegex.FindStringSubmatch(line); len(matches) > 0 {
			partNum, partTitle := matches[1], matches[2]
			partID, err = strconv.Atoi(partNum)
			sql := fmt.Sprintf(`INSERT INTO part (part_id, code_id, code, version, status, actual_date, previous_date) VALUES (%d, %d, %s, 1, TRUE, NOW(), NULL) RETURNING part_id;`, partID, codeID, partTitle)
			sqlStatements = append(sqlStatements, sql)
		} else if matches := processing.SectionRegex.FindStringSubmatch(line); len(matches) > 0 {
			sectionNum, sectionTitle := matches[1], matches[2]
			sectionID, err = strconv.Atoi(sectionNum)
			sql := fmt.Sprintf(`INSERT INTO sections (section_number, code_id, part_id, section, version, status, actual_date, previous_date) VALUES (%d, %d, %d, %s, 1, TRUE, NOW(), NULL) RETURNING section_number;`, sectionID, codeID, partID, sectionTitle)
			sqlStatements = append(sqlStatements, sql)
		} else if matches := processing.ParagraphRegex.FindStringSubmatch(line); len(matches) > 0 {
			paragraphNum, paragraphTitle := matches[1], matches[2]
			paragraphID, err = strconv.Atoi(paragraphNum)
			sql := fmt.Sprintf(`INSERT INTO paragraphs (paragraph_id, code_id, part_id, section_number, paragraph, version, status, actual_date, previous_date) VALUES (%d, %d, %d, %d, %s, 1, TRUE, NOW(), NULL) RETURNING paragraph_id;`, paragraphID, codeID, partID, sectionID, paragraphTitle)
			sqlStatements = append(sqlStatements, sql)
		} else if matches := processing.ArticleRegex.FindStringSubmatch(line); len(matches) > 0 {
			articleNum, articleTitle := matches[1], matches[2]
			articleID, err = strconv.Atoi(articleNum)
			sql := fmt.Sprintf(`INSERT INTO articles (article_number, code_id, part_id, section_number, paragraph_id, article, version, status, actual_date, previous_date) VALUES (%d, %d, %d, %d, %d, %s, 1, TRUE, NOW(), NULL) RETURNING article_number;`, articleID, codeID, partID, sectionID, paragraphID, articleTitle)
			sqlStatements = append(sqlStatements, sql)
		} else if matches := processing.ClauseRegex.FindStringSubmatch(line); len(matches) > 0 {
			clauseNum, clauseTitle := matches[1], matches[2]
			clauseID, err = strconv.Atoi(clauseNum)
			sql := fmt.Sprintf(`INSERT INTO clauses (clause_number, code_id, part_id, section_number, paragraph_id, article_number, clause, version, status, actual_date, previous_date) VALUES (%d, %d, %d, %d, %d, %d, %s, 1, TRUE, NOW(), NULL) RETURNING clause_number;`, clauseID, codeID, partID, sectionID, paragraphID, articleID, clauseTitle)
			sqlStatements = append(sqlStatements, sql)
		} else if matches := processing.SubClauseRegex.FindStringSubmatch(line); len(matches) > 0 {
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
