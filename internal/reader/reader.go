package reader

import (
	"bytes"

	"github.com/ledongthuc/pdf"
)

func ReadPDF(path string) (string, error) {
	f, r, err := pdf.Open(path)

	defer f.Close()
	if err != nil {
		return "", err
	}
	var buf bytes.Buffer
	b, err := r.GetPlainText()
	if err != nil {
		return "", err
	}
	buf.ReadFrom(b)
	return buf.String(), nil
}
