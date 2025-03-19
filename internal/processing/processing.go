package processing

import (
	"strings"
	"unicode"
)

func SplitSentences(text string) []string {
	var sentences []string
	var buffer strings.Builder

	for i := 0; i < len(text); i++ {
		buffer.WriteByte(text[i])

		if text[i] == '.' || text[i] == '?' || text[i] == '!' {

			if i > 0 && unicode.IsDigit(rune(text[i-1])) {
				continue
			}

			if i+1 < len(text) && unicode.IsLower(rune(text[i+1])) {
				continue
			}

			sentences = append(sentences, strings.TrimSpace(buffer.String()))
			buffer.Reset()
		}
	}

	if buffer.Len() > 0 {
		sentences = append(sentences, strings.TrimSpace(buffer.String()))
	}

	return sentences
}
