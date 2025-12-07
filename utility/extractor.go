package extractor

import (
	"bytes"
	"fmt"
	"io"

	"code.sajari.com/docconv"
)

// ExtractText extracts text from PDF, DOCX, or TXT files
func ExtractText(file io.Reader, mimeType string) (string, error) {
	// Read file content into buffer
	buf := new(bytes.Buffer)
	_, err := io.Copy(buf, file)
	if err != nil {
		return "", fmt.Errorf("failed to read file: %w", err)
	}

	// For plain text files, return content directly
	if mimeType == "text/plain" {
		return buf.String(), nil
	}

	// Convert document to text
	result, err := docconv.Convert(buf, mimeType, false)
	if err != nil {
		return "", fmt.Errorf("failed to extract text: %w", err)
	}

	return result.Body, nil
}
