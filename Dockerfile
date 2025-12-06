FROM golang:1.23-alpine AS builder

WORKDIR /app

# Install dependencies including build tools for docconv
RUN apk add --no-cache git gcc musl-dev poppler-utils antiword unrtf tesseract-ocr

# Copy go mod files
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Build the application with CGO enabled for docconv
RUN CGO_ENABLED=1 GOOS=linux go build -o main .

# Final stage
FROM alpine:latest

# Install runtime dependencies for document processing
RUN apk --no-cache add ca-certificates poppler-utils antiword unrtf tesseract-ocr

WORKDIR /app

# Copy the binary
COPY --from=builder /app/main .

# Expose port
EXPOSE 8080

# Run the application
CMD ["./main"]
