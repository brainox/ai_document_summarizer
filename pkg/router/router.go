package router

import (
	database "ai_document_summarizer/db"
	"ai_document_summarizer/external/external_models"
	"ai_document_summarizer/internal/config"
	"ai_document_summarizer/pkg/controller"
	"ai_document_summarizer/pkg/repository"
	storage "ai_document_summarizer/services"
	"log"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine, cfg *config.Config) {
	// Connect to database
	if err := database.Connect(cfg.DatabaseURL); err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Run migrations
	if err := database.Migrate(); err != nil {
		log.Fatalf("Failed to run migrations: %v", err)
	}

	// Initialize services
	storageClient, err := storage.NewS3Client(
		cfg.S3Endpoint,
		cfg.S3AccessKey,
		cfg.S3SecretKey,
		cfg.S3Bucket,
		cfg.S3Region,
		cfg.S3UseSSL,
	)
	if err != nil {
		log.Fatalf("Failed to initialize S3 client: %v", err)
	}

	openRouterClient := external_models.NewOpenRouterClient(cfg.OpenRouterAPIKey, cfg.OpenRouterModel)

	// Initialize repository and controller
	docRepo := repository.NewDocumentRepository()
	docController := controller.NewDocumentController(docRepo, storageClient, openRouterClient)

	// Health check
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	// Document routes
	api := r.Group("/documents")
	{
		api.POST("/upload", docController.UploadDocument)
		api.POST("/:id/analyze", docController.AnalyzeDocument)
		api.GET("/:id", docController.GetDocument)
	}
}
