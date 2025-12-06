# Setup Guide

## Prerequisites

Before you begin, ensure you have the following installed:

1. **Go 1.23+**: [Download Go](https://golang.org/dl/)
2. **Docker & Docker Compose**: [Download Docker](https://www.docker.com/products/docker-desktop)
3. **PostgreSQL** (if running locally without Docker)
4. **Minio** (if running locally without Docker)

## Step-by-Step Setup

### 1. Clone the Repository

```bash
cd /path/to/your/projects
git clone <your-repo-url>
cd ai_document_summarizer
```

### 2. Install Dependencies

```bash
go mod download
```

### 3. Get OpenRouter API Key

1. Visit [OpenRouter](https://openrouter.ai/)
2. Sign up for a free account
3. Navigate to API Keys section
4. Create a new API key
5. Copy the key (you'll need it in the next step)

### 4. Configure Environment Variables

```bash
cp .env.example .env
```

Edit `.env` and add your OpenRouter API key:

```env
OPENROUTER_API_KEY=sk-or-v1-your-key-here
```

### 5. Choose Your Setup Method

#### Option A: Using Docker (Recommended)

This is the easiest way to get started. Docker Compose will set up PostgreSQL, Minio, and the application.

```bash
# Start all services
make docker-up

# Or manually:
docker-compose up -d

# View logs
docker-compose logs -f

# Stop all services
docker-compose down
```

Services will be available at:
- **API**: http://localhost:8080
- **Minio Console**: http://localhost:9001 (credentials: minioadmin/minioadmin)
- **PostgreSQL**: localhost:5432

#### Option B: Running Locally

If you prefer to run the application locally without Docker:

**Start PostgreSQL:**
```bash
# Using Docker for PostgreSQL only
docker run -d \
  --name doc_summarizer_db \
  -e POSTGRES_USER=postgres \
  -e POSTGRES_PASSWORD=postgres \
  -e POSTGRES_DB=document_summarizer \
  -p 5432:5432 \
  postgres:15-alpine

# Or use your local PostgreSQL installation
# Create database: CREATE DATABASE document_summarizer;
```

**Start Minio:**
```bash
# Using Docker for Minio only
docker run -d \
  --name doc_summarizer_minio \
  -p 9000:9000 \
  -p 9001:9001 \
  -e MINIO_ROOT_USER=minioadmin \
  -e MINIO_ROOT_PASSWORD=minioadmin \
  minio/minio server /data --console-address ":9001"
```

**Run the Application:**
```bash
# Build and run
make run

# Or manually:
go run main.go
```

### 6. Verify Installation

Test the health endpoint:

```bash
curl http://localhost:8080/health
```

Expected response:
```json
{"status": "ok"}
```

## Usage Examples

### Upload a Document

```bash
curl -X POST http://localhost:8080/documents/upload \
  -F "file=@/path/to/your/document.pdf"
```

### Analyze Document

```bash
# Replace {id} with the document ID from upload response
curl -X POST http://localhost:8080/documents/{id}/analyze
```

### Get Document Details

```bash
curl http://localhost:8080/documents/{id}
```

See [API_EXAMPLES.md](API_EXAMPLES.md) for more detailed examples.

## Development

### Running with Auto-Reload

Install Air for automatic reloading during development:

```bash
go install github.com/air-verse/air@latest
```

Run with auto-reload:

```bash
make dev
# Or: air
```

### Available Make Commands

```bash
make help          # Show all available commands
make build         # Build the application
make run           # Run locally
make test          # Run tests
make docker-up     # Start Docker services
make docker-down   # Stop Docker services
make docker-logs   # View Docker logs
make deps          # Download dependencies
make lint          # Run linter
```

## Testing the API

Run the automated test script:

```bash
./test_api.sh
```

Or test manually using curl commands from [API_EXAMPLES.md](API_EXAMPLES.md).

## Troubleshooting

### Database Connection Issues

If you see database connection errors:

1. Verify PostgreSQL is running:
   ```bash
   docker ps | grep postgres
   ```

2. Check the DATABASE_URL in your `.env` file

3. Ensure the database exists:
   ```bash
   docker exec -it doc_summarizer_db psql -U postgres -c "\\l"
   ```

### S3/Minio Connection Issues

1. Verify Minio is running:
   ```bash
   docker ps | grep minio
   ```

2. Access Minio console at http://localhost:9001

3. Verify bucket exists or will be created automatically

### OpenRouter API Issues

If document analysis fails:

1. Verify your API key is correct in `.env`
2. Check you have credits/quota on OpenRouter
3. Try a different model (edit OPENROUTER_MODEL in `.env`)

### Text Extraction Issues

If text extraction fails:

1. Ensure the PDF/DOCX is not password-protected
2. Verify the file is not corrupted
3. Check file size (max 5MB)

### Port Already in Use

If port 8080 is already in use:

```bash
# Change PORT in .env file
PORT=8081

# Or kill the process using port 8080
lsof -ti:8080 | xargs kill -9
```

## Project Structure

```
ai_document_summarizer/
├── main.go                 # Application entry point
├── db/                     # Database setup
├── external/              
│   ├── external_models/   # External service wrappers
│   └── thirdparty/        # Third-party integrations
│       └── openrouter/    # OpenRouter API client
├── internal/
│   ├── config/            # Configuration
│   └── models/            # Data models
├── pkg/
│   ├── controller/        # HTTP request handlers
│   ├── repository/        # Database operations
│   └── router/            # Route definitions
├── services/              # Business logic
├── utility/               # Helper functions
├── .env.example           # Example environment file
├── docker-compose.yml     # Docker setup
├── Dockerfile             # Container image
├── Makefile               # Build commands
└── README.md              # Documentation
```

## Next Steps

1. **Add Authentication**: Implement JWT or API key authentication
2. **Add Rate Limiting**: Prevent abuse
3. **Add Caching**: Cache analysis results
4. **Add Webhooks**: Notify when analysis is complete
5. **Add Batch Processing**: Process multiple documents at once
6. **Add More File Types**: Support more document formats

## Support

For issues and questions:
1. Check the [troubleshooting section](#troubleshooting)
2. Review [API_EXAMPLES.md](API_EXAMPLES.md)
3. Open an issue on GitHub

## License

MIT License - see LICENSE file for details
