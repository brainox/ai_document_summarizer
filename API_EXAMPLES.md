# API Examples

## 1. Health Check

```bash
curl http://localhost:8080/health
```

## 2. Upload a Document

### Upload PDF
```bash
curl -X POST http://localhost:8080/documents/upload \
  -F "file=@/path/to/document.pdf"
```

### Upload DOCX
```bash
curl -X POST http://localhost:8080/documents/upload \
  -F "file=@/path/to/document.docx"
```

**Response:**
```json
{
  "id": "550e8400-e29b-41d4-a716-446655440000",
  "original_name": "document.pdf",
  "file_size": 245678,
  "message": "Document uploaded successfully"
}
```

## 3. Analyze Document

Replace `{document_id}` with the ID from the upload response:

```bash
curl -X POST http://localhost:8080/documents/{document_id}/analyze
```

**Response:**
```json
{
  "id": "550e8400-e29b-41d4-a716-446655440000",
  "summary": "This invoice is for software development services...",
  "document_type": "invoice",
  "metadata": {
    "date": "2024-01-15",
    "sender": "Tech Solutions Inc.",
    "recipient": "ACME Corporation",
    "amount": "$5,000.00",
    "invoice_number": "INV-2024-001"
  },
  "analyzed_at": "2024-01-15T10:30:00Z"
}
```

## 4. Get Document Details

```bash
curl http://localhost:8080/documents/{document_id}
```

**Response:**
```json
{
  "id": "550e8400-e29b-41d4-a716-446655440000",
  "original_name": "invoice.pdf",
  "file_size": 245678,
  "mime_type": "application/pdf",
  "extracted_text": "INVOICE\n\nInvoice Number: INV-2024-001...",
  "summary": "This invoice is for software development services...",
  "document_type": "invoice",
  "metadata": {
    "date": "2024-01-15",
    "sender": "Tech Solutions Inc.",
    "amount": "$5,000.00"
  },
  "analyzed_at": "2024-01-15T10:30:00Z",
  "created_at": "2024-01-15T10:00:00Z",
  "updated_at": "2024-01-15T10:30:00Z"
}
```

## Complete Workflow Example

```bash
# 1. Upload document
RESPONSE=$(curl -s -X POST http://localhost:8080/documents/upload \
  -F "file=@invoice.pdf")

# 2. Extract document ID
DOC_ID=$(echo $RESPONSE | grep -o '"id":"[^"]*"' | cut -d'"' -f4)
echo "Document ID: $DOC_ID"

# 3. Analyze document
curl -X POST http://localhost:8080/documents/$DOC_ID/analyze

# 4. Get full document details
curl http://localhost:8080/documents/$DOC_ID
```

## Using with JSON Pretty Print

```bash
# For better formatted output, pipe through jq or python
curl http://localhost:8080/documents/{document_id} | jq '.'

# Or with Python
curl http://localhost:8080/documents/{document_id} | python3 -m json.tool
```

## Error Responses

### File Too Large
```json
{
  "error": "File size exceeds 5MB limit"
}
```

### Unsupported File Type
```json
{
  "error": "Only PDF and DOCX files are supported"
}
```

### Document Not Found
```json
{
  "error": "Document not found"
}
```

### Missing API Key
```json
{
  "error": "Failed to analyze document: API returned status 401"
}
```
