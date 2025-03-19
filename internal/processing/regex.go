package processing

import (
	"github.com/dlclark/regexp2"
)

var (
	CodeRegex      = regexp2.MustCompile(`(?si)^.*?кодекс\s+республики\s+казахстан`, 0)
	PartRegex      = regexp2.MustCompile(`ЧАСТЬ\s+(\d+)\.\s+([^Р]*(?=РАЗДЕЛ|$))`, 0)
	SectionRegex   = regexp2.MustCompile(`(?si)раздел\s+(\d+)\.\s+(.*?)(?=\s*статья|\z)`, 0)
	ParagraphRegex = regexp2.MustCompile(`ПАРАГРАФ\s+(\d+)\.\s+([^С]*(?=Статья|$))`, 0)
	ArticleRegex   = regexp2.MustCompile(`Статья\s+(\d+)\.\s+([^0-9.]*)(?=\d+\.\s|$)`, 0)
	ClauseRegex    = regexp2.MustCompile(`(?m)^(\d+)\.\s+([^(]+)(?=\s+\d+\)|$)`, 0)
)
