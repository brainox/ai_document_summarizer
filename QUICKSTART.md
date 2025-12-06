# Quick Start Guide

## 🚀 Get Started in 3 Steps

### 1. Setup Environment

```bash
# Copy environment file
cp .env.example .env

# Edit .env and add your OpenRouter API key
# Get a free key at https://openrouter.ai/
```

### 2. Start Services

```bash
# Using Docker (easiest)
docker-compose up -d

# Or run locally
make run
```

### 3. Test the API

```bash
# Health check
curl http://localhost:8080/health

# Upload a document
curl -X POST http://localhost:8080/documents/upload \
  -F "file=@your-document.pdf"

# Analyze it (use the ID from upload response)
curl -X POST http://localhost:8080/documents/{id}/analyze

# Get full details
curl http://localhost:8080/documents/{id}
```

## 📚 Documentation

- **[README.md](README.md)** - Full project documentation
- **[SETUP.md](SETUP.md)** - Detailed setup instructions
- **[API_EXAMPLES.md](API_EXAMPLES.md)** - API usage examples

## 🛠️ Common Commands

```bash
make help          # Show all available commands
make docker-up     # Start all services
make docker-down   # Stop all services
make docker-logs   # View logs
make run           # Run locally
make build         # Build binary
make dev           # Run with auto-reload
```

## 🔑 Required Configuration

Only two things are required:

1. **OpenRouter API Key**: Get from https://openrouter.ai/
2. **Database**: Auto-configured with Docker Compose

## 📦 What's Included

- ✅ PDF & DOCX support (max 5MB)
- ✅ Text extraction from documents
- ✅ AI-powered analysis with OpenRouter
- ✅ Document type detection (invoice, CV, report, etc.)
- ✅ Metadata extraction (dates, amounts, sender/recipient)
- ✅ S3/Minio file storage
- ✅ PostgreSQL database
- ✅ RESTful API
- ✅ Docker support

## 🎯 API Endpoints

| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | `/health` | Health check |
| POST | `/documents/upload` | Upload PDF/DOCX |
| POST | `/documents/:id/analyze` | Analyze document with AI |
| GET | `/documents/:id` | Get document details |

## 🆘 Need Help?

1. Check [SETUP.md](SETUP.md) for troubleshooting
2. Review [API_EXAMPLES.md](API_EXAMPLES.md) for usage examples
3. Ensure `.env` has your OpenRouter API key

## 📊 Example Response

```json
{
  "id": "550e8400-e29b-41d4-a716-446655440000",
  "summary": "Invoice for software development services...",
  "document_type": "invoice",
  "metadata": {
    "date": "2024-01-15",
    "sender": "Tech Solutions Inc.",
    "amount": "$5,000.00",
    "invoice_number": "INV-2024-001"
  }
}
```

## 🌐 Service URLs

- **API**: http://localhost:8080
- **Minio Console**: http://localhost:9001 (minioadmin/minioadmin)
- **PostgreSQL**: localhost:5432

---

**Built with Go, Gin, PostgreSQL, Minio, and OpenRouter** 🚀
