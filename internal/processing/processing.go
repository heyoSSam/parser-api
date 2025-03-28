package processing

import (
	"fmt"
	"io"
	"net/http"
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

func GetDocumentText(docno string) (string, error) {
	url := fmt.Sprintf("http://91.243.71.94/api/document/details?docno=%s", docno)

	resp, err := http.Get(url)
	if err != nil {
		return "", fmt.Errorf("failed to make request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("API returned status: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read response body: %v", err)
	}

	return string(body), nil
}
