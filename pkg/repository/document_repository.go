package repository

import (
	database "ai_document_summarizer/db"
	"ai_document_summarizer/internal/models"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type DocumentRepository struct {
	db *gorm.DB
}

func NewDocumentRepository() *DocumentRepository {
	return &DocumentRepository{
		db: database.DB,
	}
}

func (r *DocumentRepository) Create(doc *models.Document) error {
	return r.db.Create(doc).Error
}

func (r *DocumentRepository) GetByID(id uuid.UUID) (*models.Document, error) {
	var doc models.Document
	err := r.db.Where("id = ?", id).First(&doc).Error
	if err != nil {
		return nil, err
	}
	return &doc, nil
}

func (r *DocumentRepository) Update(doc *models.Document) error {
	return r.db.Save(doc).Error
}

func (r *DocumentRepository) UpdateAnalysis(id uuid.UUID, summary, docType string, metadata map[string]interface{}) error {
	now := time.Now()
	return r.db.Model(&models.Document{}).
		Where("id = ?", id).
		Updates(map[string]interface{}{
			"summary":       summary,
			"document_type": docType,
			"metadata":      metadata,
			"analyzed_at":   now,
			"updated_at":    now,
		}).Error
}

func (r *DocumentRepository) Delete(id uuid.UUID) error {
	return r.db.Where("id = ?", id).Delete(&models.Document{}).Error
}
