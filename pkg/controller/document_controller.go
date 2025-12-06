package controller

import (
	"ai_document_summarizer/external/external_models"
	"ai_document_summarizer/internal/models"
	"ai_document_summarizer/pkg/repository"
	storage "ai_document_summarizer/services"
	extractor "ai_document_summarizer/utility"
	"encoding/json"
	"fmt"
	"net/http"
	"path/filepath"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type DocumentController struct {
	repo             *repository.DocumentRepository
	storageClient    *storage.S3Client
	openRouterClient *external_models.OpenRouterClient
}

func NewDocumentController(
	repo *repository.DocumentRepository,
	storageClient *storage.S3Client,
	openRouterClient *external_models.OpenRouterClient,
) *DocumentController {
	return &DocumentController{
		repo:             repo,
		storageClient:    storageClient,
		openRouterClient: openRouterClient,
	}
}

// UploadDocument handles file upload
func (dc *DocumentController) UploadDocument(c *gin.Context) {
	// Get uploaded file
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No file uploaded"})
		return
	}

	// Validate file size (max 5MB)
	const maxSize = 5 * 1024 * 1024
	if file.Size > maxSize {
		c.JSON(http.StatusBadRequest, gin.H{"error": "File size exceeds 5MB limit"})
		return
	}

	// Validate file type
	ext := strings.ToLower(filepath.Ext(file.Filename))
	if ext != ".pdf" && ext != ".docx" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Only PDF and DOCX files are supported"})
		return
	}

	// Determine MIME type
	mimeType := "application/pdf"
	if ext == ".docx" {
		mimeType = "application/vnd.openxmlformats-officedocument.wordprocessingml.document"
	}

	// Generate unique S3 key
	docID := uuid.New()
	s3Key := fmt.Sprintf("documents/%s%s", docID.String(), ext)

	// Upload to S3
	_, err = dc.storageClient.UploadFile(c.Request.Context(), file, s3Key)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to upload file to storage"})
		return
	}

	// Extract text from document
	src, err := file.Open()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read file"})
		return
	}
	defer src.Close()

	extractedText, err := extractor.ExtractText(src, mimeType)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to extract text from document"})
		return
	}

	// Create document record
	doc := &models.Document{
		ID:            docID,
		OriginalName:  file.Filename,
		S3Key:         s3Key,
		FileSize:      file.Size,
		MimeType:      mimeType,
		ExtractedText: extractedText,
	}

	if err := dc.repo.Create(doc); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save document"})
		return
	}

	c.JSON(http.StatusCreated, models.UploadResponse{
		ID:           doc.ID,
		OriginalName: doc.OriginalName,
		FileSize:     doc.FileSize,
		Message:      "Document uploaded successfully",
	})
}

// AnalyzeDocument sends document to LLM for analysis
func (dc *DocumentController) AnalyzeDocument(c *gin.Context) {
	docIDStr := c.Param("id")
	docID, err := uuid.Parse(docIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid document ID"})
		return
	}

	// Get document from database
	doc, err := dc.repo.GetByID(docID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Document not found"})
		return
	}

	// Check if document has extracted text
	if doc.ExtractedText == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Document has no extracted text"})
		return
	}

	// Send to OpenRouter for analysis
	response, err := dc.openRouterClient.AnalyzeDocument(doc.ExtractedText)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to analyze document: %v", err)})
		return
	}

	// Parse LLM response
	var analysisResult models.AnalysisResult
	if err := json.Unmarshal([]byte(response), &analysisResult); err != nil {
		// If JSON parsing fails, try to extract from markdown code blocks
		response = strings.TrimSpace(response)
		if strings.HasPrefix(response, "```json") {
			response = strings.TrimPrefix(response, "```json")
			response = strings.TrimSuffix(response, "```")
			response = strings.TrimSpace(response)
		} else if strings.HasPrefix(response, "```") {
			response = strings.TrimPrefix(response, "```")
			response = strings.TrimSuffix(response, "```")
			response = strings.TrimSpace(response)
		}

		if err := json.Unmarshal([]byte(response), &analysisResult); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse analysis result"})
			return
		}
	}

	// Update document with analysis results
	if err := dc.repo.UpdateAnalysis(docID, analysisResult.Summary, analysisResult.DocumentType, analysisResult.Metadata); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save analysis results"})
		return
	}

	c.JSON(http.StatusOK, models.AnalysisResponse{
		ID:           docID,
		Summary:      analysisResult.Summary,
		DocumentType: analysisResult.DocumentType,
		Metadata:     analysisResult.Metadata,
		AnalyzedAt:   time.Now(),
	})
}

// GetDocument retrieves document with all information
func (dc *DocumentController) GetDocument(c *gin.Context) {
	docIDStr := c.Param("id")
	docID, err := uuid.Parse(docIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid document ID"})
		return
	}

	doc, err := dc.repo.GetByID(docID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Document not found"})
		return
	}

	c.JSON(http.StatusOK, models.DocumentResponse{
		ID:            doc.ID,
		OriginalName:  doc.OriginalName,
		FileSize:      doc.FileSize,
		MimeType:      doc.MimeType,
		ExtractedText: doc.ExtractedText,
		Summary:       doc.Summary,
		DocumentType:  doc.DocumentType,
		Metadata:      doc.Metadata,
		AnalyzedAt:    doc.AnalyzedAt,
		CreatedAt:     doc.CreatedAt,
		UpdatedAt:     doc.UpdatedAt,
	})
}
