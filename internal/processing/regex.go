package processing

import (
	"github.com/dlclark/regexp2"
)

var (
	CodeRegex         = regexp2.MustCompile(`(?si)^.*?кодекс\s+республики\s+казахстан`, 0)
	PartRegex         = regexp2.MustCompile(`ЧАСТЬ\s+(\d+)\.\s+([^Р]*(?=РАЗДЕЛ|$))`, 0)
	SectionRegex      = regexp2.MustCompile(`(?si)раздел\s+(\d+)\.\s+(.*?)(?=\s*статья|\z)`, 0)
	ChapterRegex      = regexp2.MustCompile(`(?si)глава\s+(\d+)\.\s+(.*?)(?=\s*статья|\s*глава|\z)`, 0)
	ParagraphRegex    = regexp2.MustCompile(`ПАРАГРАФ\s+(\d+)\.\s+([^С]*(?=Статья|$))`, 0)
	ArticleRegex      = regexp2.MustCompile(`Статья\s+(\d+)\.\s+([^0-9.]*)(?=\d+\.\s|$)`, 0)
	ClauseRegex       = regexp2.MustCompile(`(?m)^(\d+)\.\s+([^(]+)(?=\s+\d+\)|$)`, 0)
	KAZCodeRegex      = regexp2.MustCompile(`(?si)^.*?туралы\s+кодексІ`, 0)
	KAZPartRegex      = regexp2.MustCompile(`(?si)(\d+)-бөлім\.?\s+([^Т]*(?=тарау|$))`, 0)
	KAZSectionRegex   = regexp2.MustCompile(`(?si)(\d+)-тарау\.?\s+(.*?)(?=\s*[0-9]+-бап|\z)`, 0)
	KAZChapterRegex   = regexp2.MustCompile(`(?si)(\d+)-бөлім\.?\s+(.*?)(?=\s*бап|\s*тарау|\z)`, 0)
	KAZParagraphRegex = regexp2.MustCompile(`(?si)(\d+)-параграф\.?\s+([^Б]*(?=[0-9]+-бап|$))`, 0)
	KAZArticleRegex   = regexp2.MustCompile(`(\d+)-бап\.?\s+([^0-9.]*)(?=(\d+-бап|\z))`, 0)
	KAZClauseRegex    = regexp2.MustCompile(`(?m)^(\d+)\.\s+([^(]+)(?=\s+\d+\)|$)`, 0)
)
