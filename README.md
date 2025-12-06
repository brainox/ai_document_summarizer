# AI Document Summarizer

An intelligent document processing service that extracts text from PDF and DOCX files, analyzes them using AI (via OpenRouter), and provides summaries along with metadata extraction.

## Features

- 📄 **File Upload**: Support for PDF and DOCX files (max 5MB)
- 🤖 **AI Analysis**: Automatic document summarization and classification using OpenRouter LLMs
- 🏷️ **Metadata Extraction**: Extract key information (dates, amounts, senders, etc.)
- 💾 **Storage**: S3/Minio integration for file storage
- 🗄️ **Database**: PostgreSQL for metadata and analysis results

## Tech Stack

- **Language**: Go 1.21+
- **Web Framework**: Gin
- **Database**: PostgreSQL with GORM
- **Storage**: S3/Minio
- **AI**: OpenRouter API
- **Text Extraction**: docconv

## Prerequisites

- Go 1.21 or higher
- Docker and Docker Compose (for local development)
- OpenRouter API key ([Get one here](https://openrouter.ai/))

## Quick Start

### 1. Clone and Setup

```bash
git clone <repository-url>
cd ai_document_summarizer
```

### 2. Configure Environment

Copy the example environment file and add your OpenRouter API key:

```bash
cp .env.example .env
# Edit .env and add your OPENROUTER_API_KEY
```

### 3. Install Dependencies

```bash
go mod download
```

### 4. Run with Docker Compose

```bash
docker-compose up -d
```

This will start:
- PostgreSQL database (port 5432)
- Minio object storage (port 9000, console 9001)
- Application server (port 8080)

### 5. Run Locally (without Docker)

Start PostgreSQL and Minio separately, then:

```bash
go run main.go
```

## API Endpoints

### 1. Upload Document

```bash
POST /documents/upload
Content-Type: multipart/form-data

# Example using curl
curl -X POST http://localhost:8080/documents/upload \
  -F "file=@/path/to/document.pdf"
```

**Response:**
```json
{
  "id": "uuid",
  "original_name": "document.pdf",
  "file_size": 12345,
  "message": "Document uploaded successfully"
}
```

### 2. Analyze Document

```bash
POST /documents/{id}/analyze

# Example
curl -X POST http://localhost:8080/documents/{id}/analyze
```

**Response:**
```json
{
  "id": "uuid",
  "summary": "This is a brief summary of the document...",
  "document_type": "invoice",
  "metadata": {
    "date": "2024-01-15",
    "sender": "Company Name",
    "amount": "$1,234.56"
  },
  "analyzed_at": "2024-01-15T10:30:00Z"
}
```

### 3. Get Document

```bash
GET /documents/{id}

# Example
curl http://localhost:8080/documents/{id}
```

**Response:**
```json
{
  "id": "uuid",
  "original_name": "document.pdf",
  "file_size": 12345,
  "mime_type": "application/pdf",
  "extracted_text": "Full document text...",
  "summary": "Brief summary...",
  "document_type": "invoice",
  "metadata": {
    "date": "2024-01-15",
    "amount": "$1,234.56"
  },
  "analyzed_at": "2024-01-15T10:30:00Z",
  "created_at": "2024-01-15T10:00:00Z",
  "updated_at": "2024-01-15T10:30:00Z"
}
```

### 4. Health Check

```bash
GET /health

curl http://localhost:8080/health
```

## Project Structure

```
.
├── main.go                 # Application entry point
├── db/                     # Database configuration
├── external/              
│   ├── external_models/   # External service models
│   └── thirdparty/        # Third-party integrations
│       └── openrouter/    # OpenRouter client
├── internal/
│   ├── config/            # Configuration management
│   └── models/            # Domain models
├── pkg/
│   ├── controller/        # HTTP handlers
│   ├── repository/        # Data access layer
│   └── router/            # Route definitions
├── services/              # Business logic services
├── utility/               # Helper functions
├── Dockerfile             # Docker configuration
└── docker-compose.yml     # Docker Compose setup
```

## Environment Variables

| Variable | Description | Default |
|----------|-------------|---------|
| `PORT` | Server port | `8080` |
| `DATABASE_URL` | PostgreSQL connection string | Required |
| `OPENROUTER_API_KEY` | OpenRouter API key | Required |
| `OPENROUTER_MODEL` | LLM model to use | `openai/gpt-4o-mini` |
| `S3_ENDPOINT` | S3/Minio endpoint | `localhost:9000` |
| `S3_ACCESS_KEY` | S3 access key | `minioadmin` |
| `S3_SECRET_KEY` | S3 secret key | `minioadmin` |
| `S3_BUCKET` | S3 bucket name | `documents` |
| `S3_USE_SSL` | Use SSL for S3 | `false` |
| `S3_REGION` | S3 region | `us-east-1` |

## Development

### Running Tests

```bash
go test ./...
```

### Building

```bash
go build -o ai_document_summarizer
```

## Supported Document Types

The system automatically detects and classifies:
- Invoices
- CVs/Resumes
- Reports
- Letters
- Contracts
- And more...

## Supported File Formats

- PDF (`.pdf`) - max 5MB
- DOCX (`.docx`) - max 5MB

## License

MIT

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.
