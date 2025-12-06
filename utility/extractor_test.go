package extractor

import (
	"strings"
	"testing"
)

func TestExtractText(t *testing.T) {
	tests := []struct {
		name     string
		mimeType string
		wantErr  bool
	}{
		{
			name:     "PDF mime type",
			mimeType: "application/pdf",
			wantErr:  false,
		},
		{
			name:     "DOCX mime type",
			mimeType: "application/vnd.openxmlformats-officedocument.wordprocessingml.document",
			wantErr:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a simple test reader
			reader := strings.NewReader("test content")

			_, err := ExtractText(reader, tt.mimeType)
			if (err != nil) != tt.wantErr {
				t.Errorf("ExtractText() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
