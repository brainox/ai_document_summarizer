# 🚀 Step-by-Step Guide: Run & Test the App

Follow these steps exactly to get your app running and test all endpoints.

---

## ✅ Step 1: Get Your OpenRouter API Key

1. Open your browser and go to: **https://openrouter.ai/**
2. Click **"Sign In"** (top right corner)
3. Sign up using Google, GitHub, or email
4. After signing in, click on **"Keys"** in the sidebar
5. Click **"Create Key"** button
6. Give it a name like "Document Summarizer"
7. Click **"Create"**
8. **Copy the key** (it starts with `sk-or-v1-...`)
9. Keep this tab open - you'll need this key in the next step

---

## ✅ Step 2: Configure Your Environment

Run these commands in your terminal:

```bash
# Make sure you're in the project directory
cd /Users/aguwa/Developer/HNG/ai_document_summarizer

# Copy the example environment file
cp .env.example .env

# Open the .env file in your default editor
open -e .env
```

**In the `.env` file that opens:**
1. Find the line: `OPENROUTER_API_KEY=your_openrouter_api_key_here`
2. Replace `your_openrouter_api_key_here` with your actual key
3. It should look like: `OPENROUTER_API_KEY=sk-or-v1-abc123...`
4. **Save the file** (Cmd+S)
5. Close the editor

---

## ✅ Step 3: Start Docker Services

```bash
# Start all services (PostgreSQL, Minio, and the app)
docker-compose up -d
```

**Wait about 30 seconds** for all services to start up.

**Expected output:**
```
Creating network "ai_document_summarizer_default" with the default driver
Creating volume "ai_document_summarizer_postgres_data" with default driver
Creating volume "ai_document_summarizer_minio_data" with default driver
Creating doc_summarizer_db ... done
Creating doc_summarizer_minio ... done
Creating doc_summarizer_app ... done
```

---

## ✅ Step 4: Verify Services Are Running

```bash
# Check that all containers are running
docker-compose ps
```

**Expected output** (all should show "Up"):
```
         Name                       Command               State           Ports         
---------------------------------------------------------------------------------------
doc_summarizer_app      ./main                           Up      0.0.0.0:8080->8080/tcp
doc_summarizer_db       docker-entrypoint.sh postgres    Up      0.0.0.0:5432->5432/tcp
doc_summarizer_minio    /usr/bin/docker-entrypoint...    Up      0.0.0.0:9000->9000/tcp,
                                                                  0.0.0.0:9001->9001/tcp
```

If any service shows "Exit" or "Restarting", check logs:
```bash
docker-compose logs
```

---

## ✅ Step 5: Test Health Endpoint

```bash
curl http://localhost:8080/health
```

**Expected response:**
```json
{"status":"ok"}
```

✅ If you see this, your app is running! 🎉

---

## ✅ Step 6: Create a Test Document

Let's create a sample PDF with some content to test:

```bash
# Create a simple text file
cat > test_document.txt << 'EOF'
INVOICE

Invoice Number: INV-2024-001
Date: January 15, 2024

From: Tech Solutions Inc.
123 Tech Street
San Francisco, CA 94105

To: ACME Corporation
456 Business Ave
New York, NY 10001

Services Rendered:
- Software Development (40 hours @ $125/hr): $5,000.00
- Code Review (10 hours @ $100/hr): $1,000.00

Subtotal: $6,000.00
Tax (10%): $600.00
Total Amount Due: $6,600.00

Payment Terms: Net 30
Due Date: February 15, 2024
EOF

echo "✅ Test document created: test_document.txt"
```

**Note:** For PDF testing, you can either:
- Convert this to PDF using Preview app: `File → Export as PDF`
- Or use any existing PDF/DOCX file you have

For now, let's use a tool to create a PDF:

```bash
# Install textutil if needed (comes with macOS)
textutil -convert html test_document.txt -output test_document.html

# Open in browser and save as PDF
open test_document.html
```

Then in Safari: `File → Export as PDF` → Save as `test_invoice.pdf`

**OR use an existing PDF/DOCX file you already have.**

---

## ✅ Step 7: Upload Your Document

```bash
# Upload the document (replace with your file path if different)
curl -X POST http://localhost:8080/documents/upload \
  -F "file=@test_invoice.pdf" \
  -w "\n"
```

**Expected response:**
```json
{
  "id": "550e8400-e29b-41d4-a716-446655440000",
  "original_name": "test_invoice.pdf",
  "file_size": 12345,
  "message": "Document uploaded successfully"
}
```

**✅ Copy the "id" value** - you'll need it for the next steps!

---

## ✅ Step 8: Analyze the Document

Replace `YOUR_DOCUMENT_ID` with the ID from Step 7:

```bash
# Set the document ID as a variable (paste your actual ID)
export DOC_ID="550e8400-e29b-41d4-a716-446655440000"

# Analyze the document
curl -X POST http://localhost:8080/documents/$DOC_ID/analyze \
  -w "\n"
```

**Expected response:**
```json
{
  "id": "550e8400-e29b-41d4-a716-446655440000",
  "summary": "This invoice is for software development and code review services provided by Tech Solutions Inc. to ACME Corporation, totaling $6,600.00, due by February 15, 2024.",
  "document_type": "invoice",
  "metadata": {
    "amount": "$6,600.00",
    "date": "2024-01-15",
    "due_date": "2024-02-15",
    "invoice_number": "INV-2024-001",
    "recipient": "ACME Corporation",
    "sender": "Tech Solutions Inc."
  },
  "analyzed_at": "2024-12-06T15:30:00Z"
}
```

✅ The AI has analyzed your document and extracted the metadata!

---

## ✅ Step 9: Retrieve Complete Document Data

```bash
# Get full document details
curl http://localhost:8080/documents/$DOC_ID \
  -w "\n"
```

**Expected response:**
```json
{
  "id": "550e8400-e29b-41d4-a716-446655440000",
  "original_name": "test_invoice.pdf",
  "file_size": 12345,
  "mime_type": "application/pdf",
  "extracted_text": "INVOICE\n\nInvoice Number: INV-2024-001\nDate: January 15, 2024...",
  "summary": "This invoice is for software development and code review services...",
  "document_type": "invoice",
  "metadata": {
    "amount": "$6,600.00",
    "date": "2024-01-15",
    "due_date": "2024-02-15",
    "invoice_number": "INV-2024-001",
    "recipient": "ACME Corporation",
    "sender": "Tech Solutions Inc."
  },
  "analyzed_at": "2024-12-06T15:30:00Z",
  "created_at": "2024-12-06T15:25:00Z",
  "updated_at": "2024-12-06T15:30:00Z"
}
```

✅ You now have the complete document with all analysis!

---

## ✅ Step 10: Pretty Print Results (Optional)

For better formatted JSON output:

```bash
# Install jq if you don't have it
brew install jq

# Get document with pretty formatting
curl -s http://localhost:8080/documents/$DOC_ID | jq '.'
```

---

## 🎯 Complete Test Workflow (Copy & Paste)

Here's a complete script you can run:

```bash
#!/bin/bash

echo "🚀 Testing AI Document Summarizer"
echo "=================================="
echo ""

# Test health
echo "1️⃣ Testing health endpoint..."
curl -s http://localhost:8080/health | jq '.'
echo ""

# Upload document (replace with your file)
echo "2️⃣ Uploading document..."
UPLOAD_RESPONSE=$(curl -s -X POST http://localhost:8080/documents/upload \
  -F "file=@test_invoice.pdf")
echo $UPLOAD_RESPONSE | jq '.'

# Extract document ID
DOC_ID=$(echo $UPLOAD_RESPONSE | jq -r '.id')
echo ""
echo "📄 Document ID: $DOC_ID"
echo ""

# Wait a moment
sleep 2

# Analyze document
echo "3️⃣ Analyzing document..."
curl -s -X POST http://localhost:8080/documents/$DOC_ID/analyze | jq '.'
echo ""

# Wait for analysis to complete
sleep 3

# Get full document
echo "4️⃣ Retrieving full document..."
curl -s http://localhost:8080/documents/$DOC_ID | jq '.'
echo ""

echo "✅ All tests completed!"
```

Save this as `test.sh`, make it executable, and run:

```bash
chmod +x test.sh
./test.sh
```

---

## 📊 View Running Services

### Check Application Logs
```bash
# View all logs
docker-compose logs -f

# View only app logs
docker-compose logs -f app
```

### Access Minio Console (File Storage)
1. Open browser: **http://localhost:9001**
2. Login: `minioadmin` / `minioadmin`
3. You'll see your uploaded documents in the `documents` bucket

### Access Database
```bash
# Connect to PostgreSQL
docker exec -it doc_summarizer_db psql -U postgres -d document_summarizer

# View documents table
SELECT id, original_name, document_type, created_at FROM documents;

# Exit
\q
```

---

## 🛑 Stop Services

```bash
# Stop all services
docker-compose down

# Stop and remove volumes (cleans everything)
docker-compose down -v
```

---

## 🔄 Restart Services

```bash
# Restart all services
docker-compose restart

# Rebuild and restart (if you made code changes)
docker-compose down
docker-compose up -d --build
```

---

## 🐛 Troubleshooting

### Port Already in Use
```bash
# Find what's using port 8080
lsof -ti:8080

# Kill the process
lsof -ti:8080 | xargs kill -9

# Or change port in .env file
PORT=8081
```

### OpenRouter API Errors
```bash
# Check if API key is set correctly
docker-compose exec app env | grep OPENROUTER

# View app logs for errors
docker-compose logs app
```

### Database Connection Failed
```bash
# Restart database
docker-compose restart postgres

# Check database status
docker-compose ps postgres
```

### File Upload Fails
```bash
# Check Minio is running
docker-compose ps minio

# View Minio logs
docker-compose logs minio
```

---

## ✅ Success Checklist

- [ ] All services running (`docker-compose ps`)
- [ ] Health endpoint returns `{"status":"ok"}`
- [ ] Document uploaded successfully
- [ ] Analysis returns summary and metadata
- [ ] Can retrieve full document data
- [ ] Minio console accessible at http://localhost:9001
- [ ] No errors in logs (`docker-compose logs`)

---

## 🎉 Next Steps

Now that everything is working:

1. **Try different document types**: Upload a CV, report, or contract
2. **View files in Minio**: Check http://localhost:9001
3. **Explore the API**: See `API_EXAMPLES.md` for more examples
4. **Customize**: Modify the OpenRouter prompt in `external/thirdparty/openrouter/client.go`

---

## 📞 Need Help?

- Check `SETUP.md` for troubleshooting
- View logs: `docker-compose logs -f`
- Ensure OpenRouter API key is valid
- Verify all services are "Up": `docker-compose ps`

**You're all set! Happy testing! 🚀**
