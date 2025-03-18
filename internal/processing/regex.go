package processing

import "regexp"

var (
	CodeRegex      = regexp.MustCompile(`(?s)^.*?КОДЕКС РЕСПУБЛИКИ КАЗАХСТАН`)
	PartRegex      = regexp.MustCompile(`ЧАСТЬ\s+(\d+)\.\s+(.+)`)
	SectionRegex   = regexp.MustCompile(`РАЗДЕЛ\s+(\d+)\.\s+(.+)`)
	ParagraphRegex = regexp.MustCompile(`ПАРАГРАФ\s+(\d+)\.\s+(.+)`)
	ArticleRegex   = regexp.MustCompile(`Статья\s+(\d+)\.\s+(.+)`)
	ClauseRegex    = regexp.MustCompile(`(?m)^(\d+)\.\s+(.+)$`)
	SubClauseRegex = regexp.MustCompile(`(?m)^(\d+)\)\s+(.+)$`)
)
