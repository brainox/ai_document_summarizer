.PHONY: help build run test clean docker-up docker-down migrate

help: ## Show this help message
	@echo 'Usage: make [target]'
	@echo ''
	@echo 'Available targets:'
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "  %-15s %s\n", $$1, $$2}' $(MAKEFILE_LIST)

build: ## Build the application
	go build -o bin/ai_document_summarizer main.go

run: ## Run the application locally
	go run main.go

test: ## Run tests
	go test -v ./...

clean: ## Clean build artifacts
	rm -rf bin/
	rm -f ai_document_summarizer

docker-up: ## Start all services with Docker Compose
	docker-compose up -d

docker-down: ## Stop all services
	docker-compose down

docker-logs: ## Show last 50 lines of logs from all services
	docker-compose logs --tail=50

docker-logs-follow: ## Follow logs from all services (press Ctrl+C to stop)
	docker-compose logs -f

docker-rebuild: ## Rebuild and restart all services
	docker-compose down
	docker-compose up -d --build

deps: ## Download dependencies
	go mod download
	go mod tidy

lint: ## Run linter
	go fmt ./...
	go vet ./...

dev: ## Run in development mode with auto-reload
	@which air > /dev/null || go install github.com/air-verse/air@latest
	air

# Heroku deployment commands
heroku-login: ## Login to Heroku
	heroku login

heroku-create: ## Create new Heroku app with container stack
	@read -p "Enter app name (or leave blank for random): " app_name; \
	if [ -z "$$app_name" ]; then \
		heroku create --stack container; \
	else \
		heroku create $$app_name --stack container; \
	fi

heroku-addons: ## Add PostgreSQL addon
	heroku addons:create heroku-postgresql:essential-0

heroku-config: ## Set environment variables on Heroku
	@read -p "Enter OpenRouter API key: " api_key; \
	heroku config:set OPENROUTER_API_KEY=$$api_key
	@read -p "Enter AWS Access Key ID: " aws_key; \
	heroku config:set AWS_ACCESS_KEY_ID=$$aws_key
	@read -p "Enter AWS Secret Access Key: " aws_secret; \
	heroku config:set AWS_SECRET_ACCESS_KEY=$$aws_secret
	@read -p "Enter S3 Bucket name: " bucket; \
	heroku config:set S3_BUCKET=$$bucket
	@read -p "Enter S3 Endpoint (leave empty for AWS S3): " endpoint; \
	if [ -n "$$endpoint" ]; then \
		heroku config:set S3_ENDPOINT=$$endpoint; \
	fi
	heroku config:set OPENROUTER_MODEL=openai/gpt-4o-mini
	heroku config:set AWS_REGION=us-east-1
	heroku config:set MAX_FILE_SIZE_MB=5

heroku-deploy: ## Deploy to Heroku
	git push heroku main

heroku-logs: ## View Heroku logs
	heroku logs --tail

heroku-open: ## Open app in browser
	heroku open

heroku-setup: heroku-login heroku-create heroku-addons heroku-config ## Complete Heroku setup
	@echo "✓ Heroku setup complete! Run 'make heroku-deploy' to deploy."

heroku-info: ## Show Heroku app info
	heroku info
	@echo "\nConfig variables:"
	heroku config

.DEFAULT_GOAL := help
