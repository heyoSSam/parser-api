//package processing
//
//import (
//	"regexp"
//	"strings"
//	"unicode"
//)
//
////func SplitSentences(text string) []string {
////	var sentences []string
////	var buffer strings.Builder
////
////	for i := 0; i < len(text); i++ {
////		buffer.WriteByte(text[i])
////
////		if text[i] == '.' || text[i] == '?' || text[i] == '!' {
////
////			if i > 0 && unicode.IsDigit(rune(text[i-1])) {
////				continue
////			}
////
////			if i+1 < len(text) && unicode.IsLower(rune(text[i+1])) {
////				continue
////			}
////
////			sentences = append(sentences, strings.TrimSpace(buffer.String()))
////			buffer.Reset()
////		}
////	}
////
////	if buffer.Len() > 0 {
////		sentences = append(sentences, strings.TrimSpace(buffer.String()))
////	}
////
////	return sentences
////}
//
//func SplitSentences(text string) []string {
//	re := regexp.MustCompile(`(?=(РАЗДЕЛ \d+|Глава \d+|Статья \d+))`)
//
//	// Разделяем текст
//	parts := re.Split(text, -1)
//	matches := re.FindAllString(text, -1)
//
//	var result []string
//
//	// Склеиваем совпадения с их соответствующими частями
//	for i := 1; i < len(parts); i++ {
//		result = append(result, strings.TrimSpace(matches[i-1]+" "+parts[i]))
//	}
//
//	return result
//}

package processing

import (
	"regexp"
	"strings"
	"unicode"
)

var titleRegex = regexp.MustCompile(`(?i)\b(РАЗДЕЛ|ГЛАВА|СТАТЬЯ|ПУНКТ|ЧАСТЬ)\s+\d+`)

func SplitSentences(text string) []string {
	var sentences []string
	var buffer strings.Builder
	var lastTitle string

	words := strings.Fields(text)
	wordCount := len(words)

	for i := 0; i < wordCount; i++ {
		word := words[i]

		if buffer.Len() > 0 {
			buffer.WriteString(" ")
		}

		buffer.WriteString(word)

		if titleRegex.MatchString(buffer.String()) {
			lastTitle = buffer.String()
			buffer.Reset()
			continue
		}

		if strings.HasSuffix(word, ".") || strings.HasSuffix(word, "!") {

			if i > 0 && unicode.IsDigit(rune(words[i-1][0])) {
				continue
			}

			if i+1 < wordCount && unicode.IsLower(rune(words[i+1][0])) {
				continue
			}

			if lastTitle != "" {
				sentences = append(sentences, lastTitle+" "+strings.TrimSpace(buffer.String()))
				lastTitle = ""
			} else {
				sentences = append(sentences, strings.TrimSpace(buffer.String()))
			}

			buffer.Reset()
			continue
		}
	}

	if buffer.Len() > 0 {
		if lastTitle != "" {
			sentences = append(sentences, lastTitle+" "+strings.TrimSpace(buffer.String()))
		} else {
			sentences = append(sentences, strings.TrimSpace(buffer.String()))
		}
	}

	return sentences
}
