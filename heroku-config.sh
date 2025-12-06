#!/bin/bash

# Heroku Configuration Script for AI Document Summarizer
# This script helps you set all required environment variables

APP_NAME="ai-document-summarizer"

echo "=================================="
echo "Heroku Configuration Setup"
echo "=================================="
echo ""

# OpenRouter API Key
echo "1. OpenRouter API Key"
read -p "   Enter your OpenRouter API key: " OPENROUTER_KEY
heroku config:set OPENROUTER_API_KEY="$OPENROUTER_KEY" -a $APP_NAME

# S3 Configuration
echo ""
echo "2. S3 Storage Configuration"
echo "   (You need S3-compatible storage: AWS S3, Cloudflare R2, DigitalOcean Spaces, etc.)"
read -p "   Enter AWS Access Key ID: " AWS_KEY
heroku config:set AWS_ACCESS_KEY_ID="$AWS_KEY" -a $APP_NAME

read -p "   Enter AWS Secret Access Key: " AWS_SECRET
heroku config:set AWS_SECRET_ACCESS_KEY="$AWS_SECRET" -a $APP_NAME

read -p "   Enter S3 Bucket Name: " S3_BUCKET
heroku config:set S3_BUCKET="$S3_BUCKET" -a $APP_NAME

echo ""
read -p "   Are you using AWS S3 or a different provider? (aws/other): " S3_PROVIDER
if [ "$S3_PROVIDER" != "aws" ]; then
    read -p "   Enter S3 Endpoint URL (e.g., https://xxx.r2.cloudflarestorage.com): " S3_ENDPOINT
    heroku config:set S3_ENDPOINT="$S3_ENDPOINT" -a $APP_NAME
    heroku config:set AWS_REGION="auto" -a $APP_NAME
else
    read -p "   Enter AWS Region (default: us-east-1): " AWS_REGION
    AWS_REGION=${AWS_REGION:-us-east-1}
    heroku config:set AWS_REGION="$AWS_REGION" -a $APP_NAME
fi

# Optional Settings
echo ""
echo "3. Optional Settings (using defaults)"
heroku config:set OPENROUTER_MODEL="openai/gpt-4o-mini" -a $APP_NAME
heroku config:set MAX_FILE_SIZE_MB="5" -a $APP_NAME

echo ""
echo "=================================="
echo "Configuration Complete!"
echo "=================================="
echo ""
echo "Your app configuration:"
heroku config -a $APP_NAME

echo ""
echo "Next step: Deploy your app with 'git push heroku main'"
