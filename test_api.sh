#!/bin/bash

echo "Testing AI Document Summarizer API"
echo "==================================="
echo ""

BASE_URL="http://localhost:8080"

# Color codes
GREEN='\033[0;32m'
RED='\033[0;31m'
NC='\033[0m' # No Color

# Test 1: Health Check
echo "1. Testing Health Check..."
response=$(curl -s -o /dev/null -w "%{http_code}" "$BASE_URL/health")
if [ "$response" -eq 200 ]; then
    echo -e "${GREEN}✓ Health check passed${NC}"
else
    echo -e "${RED}✗ Health check failed (HTTP $response)${NC}"
fi
echo ""

# Find a test PDF file
TEST_FILE="sample.pdf"

# Test 2: Upload Document
echo "2. Testing Document Upload..."
if [ -n "$TEST_FILE" ]; then
    echo "Using file: $TEST_FILE"
    response=$(curl -s -X POST "$BASE_URL/documents/upload" \
        -F "file=@$TEST_FILE")
    
    http_code=$(curl -s -o /dev/null -w "%{http_code}" -X POST "$BASE_URL/documents/upload" \
        -F "file=@$TEST_FILE")
    body="$response"
    
    if [ "$http_code" -eq 201 ]; then
        echo -e "${GREEN}✓ Document upload passed${NC}"
        DOC_ID=$(echo "$body" | python3 -c "import sys, json; print(json.load(sys.stdin)['id'])" 2>/dev/null)
        
        if [ -z "$DOC_ID" ]; then
            # Fallback to grep if python fails
            DOC_ID=$(echo "$body" | grep -o '"id":"[^"]*"' | cut -d'"' -f4)
        fi
        
        echo "Document ID: $DOC_ID"
        
        if [ -n "$DOC_ID" ]; then
            # Test 3: Analyze Document
            echo ""
            echo "3. Testing Document Analysis..."
            sleep 2
            
            analyze_response=$(curl -s -X POST "$BASE_URL/documents/$DOC_ID/analyze")
            analyze_code=$(curl -s -o /dev/null -w "%{http_code}" -X POST "$BASE_URL/documents/$DOC_ID/analyze")
            
            if [ "$analyze_code" -eq 200 ]; then
                echo -e "${GREEN}✓ Document analysis passed${NC}"
                echo "$analyze_response" | python3 -m json.tool 2>/dev/null || echo "$analyze_response"
            else
                echo -e "${RED}✗ Document analysis failed (HTTP $analyze_code)${NC}"
                echo "$analyze_response"
            fi
            
            # Test 4: Get Document
            echo ""
            echo "4. Testing Get Document..."
            sleep 1
            
            get_response=$(curl -s "$BASE_URL/documents/$DOC_ID")
            get_code=$(curl -s -o /dev/null -w "%{http_code}" "$BASE_URL/documents/$DOC_ID")
            
            if [ "$get_code" -eq 200 ]; then
                echo -e "${GREEN}✓ Get document passed${NC}"
                # Show only summary, not full extracted text
                echo "$get_response" | python3 -c "import sys, json; d=json.load(sys.stdin); del d['extracted_text']; print(json.dumps(d, indent=2))" 2>/dev/null || echo "$get_response"
            else
                echo -e "${RED}✗ Get document failed (HTTP $get_code)${NC}"
                echo "$get_response"
            fi
        else
            echo -e "${RED}✗ Could not extract document ID${NC}"
        fi
    else
        echo -e "${RED}✗ Document upload failed (HTTP $http_code)${NC}"
        echo "$body"
    fi
else
    echo -e "${RED}✗ No test PDF found. Please add sample.pdf, HelloBackendEngineers.pdf, or iOS_Prep_pdf.pdf${NC}"
fi

echo ""
echo "==================================="
echo "Testing completed"
