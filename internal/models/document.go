package models

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"time"

	"github.com/google/uuid"
)

type JSONB map[string]interface{}

func (j JSONB) Value() (driver.Value, error) {
	return json.Marshal(j)
}

func (j *JSONB) Scan(value interface{}) error {
	if value == nil {
		*j = nil
		return nil
	}
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}
	return json.Unmarshal(bytes, j)
}

type Document struct {
	ID            uuid.UUID  `json:"id" gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	OriginalName  string     `json:"original_name" gorm:"not null"`
	S3Key         string     `json:"s3_key" gorm:"not null"`
	FileSize      int64      `json:"file_size"`
	MimeType      string     `json:"mime_type"`
	ExtractedText string     `json:"extracted_text" gorm:"type:text"`
	Summary       string     `json:"summary,omitempty" gorm:"type:text"`
	DocumentType  string     `json:"document_type,omitempty"`
	Metadata      JSONB      `json:"metadata,omitempty" gorm:"type:jsonb"`
	AnalyzedAt    *time.Time `json:"analyzed_at,omitempty"`
	CreatedAt     time.Time  `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt     time.Time  `json:"updated_at" gorm:"autoUpdateTime"`
}

type UploadResponse struct {
	ID           uuid.UUID `json:"id"`
	OriginalName string    `json:"original_name"`
	FileSize     int64     `json:"file_size"`
	Message      string    `json:"message"`
}

type AnalysisResult struct {
	Summary      string                 `json:"summary"`
	DocumentType string                 `json:"document_type"`
	Metadata     map[string]interface{} `json:"metadata"`
}

type AnalysisResponse struct {
	ID           uuid.UUID              `json:"id"`
	Summary      string                 `json:"summary"`
	DocumentType string                 `json:"document_type"`
	Metadata     map[string]interface{} `json:"metadata"`
	AnalyzedAt   time.Time              `json:"analyzed_at"`
}

type DocumentResponse struct {
	ID            uuid.UUID              `json:"id"`
	OriginalName  string                 `json:"original_name"`
	FileSize      int64                  `json:"file_size"`
	MimeType      string                 `json:"mime_type"`
	ExtractedText string                 `json:"extracted_text"`
	Summary       string                 `json:"summary,omitempty"`
	DocumentType  string                 `json:"document_type,omitempty"`
	Metadata      map[string]interface{} `json:"metadata,omitempty"`
	AnalyzedAt    *time.Time             `json:"analyzed_at,omitempty"`
	CreatedAt     time.Time              `json:"created_at"`
	UpdatedAt     time.Time              `json:"updated_at"`
}
