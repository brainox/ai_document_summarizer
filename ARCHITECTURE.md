# Architecture Overview

## System Architecture

```
┌─────────────────────────────────────────────────────────────┐
│                         Client                               │
│                    (curl, Postman, etc)                      │
└─────────────────────────┬───────────────────────────────────┘
                          │ HTTP/REST
                          ▼
┌─────────────────────────────────────────────────────────────┐
│                      Gin Web Server                          │
│                      (Port 8080)                             │
└─────────────────────────┬───────────────────────────────────┘
                          │
        ┌─────────────────┼─────────────────┐
        │                 │                 │
        ▼                 ▼                 ▼
┌──────────────┐  ┌──────────────┐  ┌──────────────┐
│   Upload     │  │   Analyze    │  │  Retrieve    │
│  Controller  │  │  Controller  │  │  Controller  │
└──────┬───────┘  └──────┬───────┘  └──────┬───────┘
       │                 │                 │
       ▼                 ▼                 ▼
┌──────────────────────────────────────────────────┐
│              Document Repository                  │
│            (Database Operations)                  │
└──────────────────┬───────────────────────────────┘
                   │
     ┌─────────────┼─────────────────┐
     │             │                 │
     ▼             ▼                 ▼
┌─────────┐  ┌──────────┐    ┌──────────────┐
│PostgreSQL│ │  Minio   │    │  OpenRouter  │
│   DB     │ │  (S3)    │    │     API      │
└──────────┘ └──────────┘    └──────────────┘
```

## Component Breakdown

### 1. Entry Point (`main.go`)
- Initializes configuration
- Sets up Gin router
- Starts HTTP server

### 2. Configuration (`internal/config/`)
- Loads environment variables
- Provides configuration to all components
- Handles defaults

### 3. Database Layer (`db/`)
- **Connection**: Establishes PostgreSQL connection
- **Migration**: Auto-creates tables using GORM
- **Models**: Defines document structure

### 4. Data Models (`internal/models/`)
- **Document**: Core document entity
- **Request/Response**: API DTOs
- **Analysis Result**: LLM response structure

### 5. Repository Layer (`pkg/repository/`)
- **Abstraction**: Database operations abstraction
- **CRUD**: Create, Read, Update, Delete operations
- **Queries**: Specialized queries for document management

### 6. Controller Layer (`pkg/controller/`)
- **HTTP Handlers**: Process HTTP requests
- **Validation**: Input validation
- **Response Formatting**: JSON response generation
- **Error Handling**: Centralized error handling

### 7. Services Layer (`services/`)
- **Storage Service**: S3/Minio file operations
  - Upload files
  - Retrieve files
  - Delete files
  - Bucket management

### 8. External Services (`external/`)
- **OpenRouter Client**: LLM integration
  - Sends document text to AI
  - Parses analysis results
  - Error handling for API calls

### 9. Utilities (`utility/`)
- **Text Extractor**: 
  - PDF text extraction
  - DOCX text extraction
  - Uses docconv library

### 10. Router (`pkg/router/`)
- **Route Definition**: Maps URLs to handlers
- **Middleware**: (Future: auth, logging, CORS)
- **Dependency Injection**: Wires up all components

## Data Flow

### Upload Flow

```
1. Client uploads PDF/DOCX file
   ↓
2. Controller validates file (size, type)
   ↓
3. File uploaded to S3/Minio
   ↓
4. Text extracted from document
   ↓
5. Document record created in database
   ↓
6. Response returned with document ID
```

### Analysis Flow

```
1. Client requests analysis with document ID
   ↓
2. Controller retrieves document from database
   ↓
3. Extracted text sent to OpenRouter API
   ↓
4. LLM analyzes and returns:
   - Summary
   - Document type
   - Metadata
   ↓
5. Results saved to database
   ↓
6. Response returned with analysis results
```

### Retrieval Flow

```
1. Client requests document by ID
   ↓
2. Controller retrieves from database
   ↓
3. Full document data returned including:
   - File metadata
   - Extracted text
   - Analysis results
   - Timestamps
```

## Technology Stack

### Backend Framework
- **Gin**: Fast HTTP web framework
- **GORM**: ORM for database operations
- **UUID**: Unique identifiers

### Storage
- **PostgreSQL**: Relational database for metadata
- **Minio**: S3-compatible object storage for files

### External Services
- **OpenRouter**: LLM API aggregator
- **docconv**: Document conversion library

### Development Tools
- **Air**: Hot reload for development
- **Docker Compose**: Local development environment
- **Make**: Build automation

## Database Schema

### documents Table

```sql
CREATE TABLE documents (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    original_name VARCHAR(255) NOT NULL,
    s3_key VARCHAR(500) NOT NULL,
    file_size BIGINT,
    mime_type VARCHAR(100),
    extracted_text TEXT,
    summary TEXT,
    document_type VARCHAR(100),
    metadata JSONB,
    analyzed_at TIMESTAMP,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_documents_created_at ON documents(created_at);
CREATE INDEX idx_documents_document_type ON documents(document_type);
```

## API Design

### RESTful Principles
- Standard HTTP methods (GET, POST)
- Resource-based URLs
- JSON payloads
- HTTP status codes

### Endpoints

| Endpoint | Method | Purpose | Request | Response |
|----------|--------|---------|---------|----------|
| `/health` | GET | Health check | - | `{status: "ok"}` |
| `/documents/upload` | POST | Upload file | multipart/form-data | Document metadata |
| `/documents/:id/analyze` | POST | Analyze document | - | Analysis results |
| `/documents/:id` | GET | Get document | - | Full document data |

## Security Considerations

### Current Implementation
- File size validation (max 5MB)
- File type validation (PDF, DOCX only)
- UUID for unpredictable IDs

### Recommended Additions
1. **Authentication**: JWT or API keys
2. **Rate Limiting**: Prevent abuse
3. **CORS**: Configure allowed origins
4. **File Scanning**: Malware detection
5. **Encryption**: Encrypt files at rest
6. **HTTPS**: TLS/SSL in production

## Scalability Considerations

### Current Limitations
- Single server instance
- Synchronous processing
- No caching

### Scaling Strategies

#### Horizontal Scaling
- Deploy multiple app instances
- Load balancer in front
- Shared database and storage

#### Asynchronous Processing
```
┌──────────┐    ┌──────────┐    ┌──────────┐
│  Upload  │───▶│  Queue   │───▶│  Worker  │
│  API     │    │ (Redis)  │    │  Pool    │
└──────────┘    └──────────┘    └──────────┘
                                      │
                                      ▼
                                ┌──────────┐
                                │OpenRouter│
                                └──────────┘
```

#### Caching
- Redis for frequently accessed documents
- CDN for file downloads
- Cache analysis results

#### Database Optimization
- Read replicas for scaling reads
- Connection pooling
- Indexing on common queries

## Error Handling Strategy

### Levels
1. **Validation Errors**: 400 Bad Request
2. **Not Found**: 404 Not Found
3. **Server Errors**: 500 Internal Server Error
4. **External API Errors**: 502 Bad Gateway

### Error Response Format
```json
{
  "error": "Descriptive error message"
}
```

## Monitoring & Observability

### Recommended Additions
1. **Logging**: Structured logging with levels
2. **Metrics**: Prometheus + Grafana
3. **Tracing**: OpenTelemetry
4. **Alerting**: PagerDuty, Slack notifications

### Key Metrics to Track
- Request rate
- Response time
- Error rate
- Document processing time
- LLM API latency
- Storage usage

## Future Enhancements

### Features
1. Batch processing
2. Webhook notifications
3. Multiple LLM providers
4. Custom extraction templates
5. OCR for scanned documents
6. Multi-language support

### Technical
1. GraphQL API
2. WebSocket for real-time updates
3. Microservices architecture
4. Event-driven architecture
5. Machine learning models
6. Advanced search capabilities

## Development Workflow

### Local Development
```bash
# Install dependencies
make deps

# Run with hot reload
make dev

# Run tests
make test

# Build
make build
```

### Docker Development
```bash
# Start all services
make docker-up

# View logs
make docker-logs

# Rebuild
make docker-rebuild

# Stop services
make docker-down
```

### Testing Strategy
1. **Unit Tests**: Test individual functions
2. **Integration Tests**: Test component interactions
3. **E2E Tests**: Test complete workflows
4. **Load Tests**: Test performance under load

## Deployment

### Docker Deployment
```bash
# Build image
docker build -t ai-doc-summarizer .

# Run container
docker run -d \
  -p 8080:8080 \
  -e DATABASE_URL=... \
  -e OPENROUTER_API_KEY=... \
  ai-doc-summarizer
```

### Cloud Deployment Options
- **AWS**: ECS, RDS, S3
- **Google Cloud**: Cloud Run, Cloud SQL, GCS
- **Azure**: Container Instances, PostgreSQL, Blob Storage
- **Kubernetes**: Deployments, Services, PVCs

## License

MIT License - See LICENSE file for details.
