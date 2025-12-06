# Document Types & Metadata Examples

This document shows examples of different document types and the metadata that will be extracted.

## 📄 Invoice

### Detected Type
`invoice`

### Typical Metadata
```json
{
  "date": "2024-01-15",
  "invoice_number": "INV-2024-001",
  "sender": "Tech Solutions Inc.",
  "recipient": "ACME Corporation",
  "amount": "$5,000.00",
  "due_date": "2024-02-15",
  "payment_terms": "Net 30",
  "items": "Software development services"
}
```

### Summary Example
```
This invoice is for software development services provided by Tech Solutions Inc. 
to ACME Corporation for the amount of $5,000.00, due by February 15, 2024.
```

---

## 📝 CV/Resume

### Detected Type
`cv` or `resume`

### Typical Metadata
```json
{
  "name": "John Doe",
  "email": "john.doe@email.com",
  "phone": "+1-234-567-8900",
  "position": "Senior Software Engineer",
  "experience_years": "8",
  "education": "B.S. Computer Science",
  "skills": ["Python", "Go", "JavaScript", "AWS"],
  "last_employer": "Tech Corp"
}
```

### Summary Example
```
Resume of John Doe, a Senior Software Engineer with 8 years of experience in 
software development, specializing in Python, Go, and cloud technologies.
```

---

## 📊 Report

### Detected Type
`report`

### Typical Metadata
```json
{
  "title": "Q4 2024 Financial Report",
  "date": "2024-01-05",
  "author": "Finance Department",
  "department": "Finance",
  "report_type": "Financial",
  "period": "Q4 2024",
  "key_figures": {
    "revenue": "$2.5M",
    "profit": "$500K",
    "growth": "15%"
  }
}
```

### Summary Example
```
Financial report for Q4 2024 showing revenue of $2.5M with 15% growth compared 
to the previous quarter, and a profit margin of $500K.
```

---

## 💼 Contract

### Detected Type
`contract`

### Typical Metadata
```json
{
  "contract_type": "Employment Agreement",
  "date": "2024-01-01",
  "party_a": "ABC Company Inc.",
  "party_b": "Jane Smith",
  "start_date": "2024-02-01",
  "duration": "Indefinite",
  "salary": "$120,000 per annum",
  "notice_period": "30 days"
}
```

### Summary Example
```
Employment agreement between ABC Company Inc. and Jane Smith, effective 
February 1, 2024, with an annual salary of $120,000 and 30-day notice period.
```

---

## ✉️ Letter

### Detected Type
`letter`

### Typical Metadata
```json
{
  "date": "2024-01-15",
  "sender": "Human Resources Department",
  "recipient": "All Employees",
  "subject": "Policy Update",
  "letter_type": "Business Correspondence",
  "urgency": "Normal"
}
```

### Summary Example
```
Business letter from HR department to all employees regarding updated company 
policies, dated January 15, 2024.
```

---

## 📋 Form

### Detected Type
`form`

### Typical Metadata
```json
{
  "form_type": "Application Form",
  "form_number": "F-2024-001",
  "date_submitted": "2024-01-15",
  "applicant": "John Smith",
  "status": "Pending",
  "department": "Human Resources"
}
```

### Summary Example
```
Application form F-2024-001 submitted by John Smith on January 15, 2024, 
currently pending review by the HR department.
```

---

## 📖 Manual/Guide

### Detected Type
`manual` or `guide`

### Typical Metadata
```json
{
  "title": "User Manual - Software Installation",
  "version": "2.0",
  "date": "2024-01-01",
  "author": "Documentation Team",
  "pages": "25",
  "product": "Enterprise Software Suite"
}
```

### Summary Example
```
User manual version 2.0 for Enterprise Software Suite installation, providing 
step-by-step instructions for system setup and configuration.
```

---

## 📑 Proposal

### Detected Type
`proposal`

### Typical Metadata
```json
{
  "title": "Website Redesign Proposal",
  "date": "2024-01-15",
  "submitted_by": "Design Agency XYZ",
  "submitted_to": "ABC Corporation",
  "project_value": "$50,000",
  "timeline": "3 months",
  "valid_until": "2024-02-15"
}
```

### Summary Example
```
Business proposal from Design Agency XYZ to ABC Corporation for website 
redesign project valued at $50,000 with 3-month timeline.
```

---

## 📧 Memo

### Detected Type
`memo` or `memorandum`

### Typical Metadata
```json
{
  "date": "2024-01-15",
  "from": "Management Team",
  "to": "All Staff",
  "subject": "Office Closure Notice",
  "priority": "High",
  "reference_number": "MEM-2024-003"
}
```

### Summary Example
```
Internal memo from Management Team informing all staff about upcoming office 
closure dates and remote work arrangements.
```

---

## 🏥 Medical Document

### Detected Type
`medical_document` or `medical_report`

### Typical Metadata
```json
{
  "document_type": "Lab Report",
  "patient_id": "P-123456",
  "date": "2024-01-15",
  "doctor": "Dr. Sarah Johnson",
  "facility": "City Medical Center",
  "test_type": "Blood Test"
}
```

### Summary Example
```
Laboratory test report for patient P-123456, conducted by Dr. Sarah Johnson 
at City Medical Center on January 15, 2024.
```

---

## 🎓 Academic Document

### Detected Type
`academic` or `thesis` or `research_paper`

### Typical Metadata
```json
{
  "title": "Machine Learning Applications in Healthcare",
  "author": "Dr. Emily Chen",
  "institution": "State University",
  "date": "2024-01-01",
  "document_type": "Research Paper",
  "pages": "45",
  "field": "Computer Science"
}
```

### Summary Example
```
Research paper on machine learning applications in healthcare by Dr. Emily Chen 
from State University, exploring AI-driven diagnosis systems.
```

---

## 💳 Receipt

### Detected Type
`receipt`

### Typical Metadata
```json
{
  "date": "2024-01-15",
  "merchant": "Office Supplies Inc.",
  "amount": "$245.67",
  "payment_method": "Credit Card",
  "receipt_number": "R-2024-001234",
  "items": "Office supplies and stationery"
}
```

### Summary Example
```
Purchase receipt from Office Supplies Inc. for $245.67 paid by credit card, 
receipt number R-2024-001234, dated January 15, 2024.
```

---

## 📈 Presentation

### Detected Type
`presentation`

### Typical Metadata
```json
{
  "title": "Q4 Marketing Strategy",
  "presenter": "Marketing Team",
  "date": "2024-01-15",
  "audience": "Executive Board",
  "slides": "30",
  "topic": "Marketing Strategy"
}
```

### Summary Example
```
Marketing presentation covering Q4 strategy, presented to Executive Board 
by the Marketing Team, focusing on digital campaigns and growth targets.
```

---

## 🔧 Technical Documentation

### Detected Type
`technical_documentation`

### Typical Metadata
```json
{
  "title": "API Integration Guide",
  "version": "3.1",
  "date": "2024-01-01",
  "author": "Engineering Team",
  "technology": "REST API",
  "audience": "Developers"
}
```

### Summary Example
```
Technical documentation for REST API integration, version 3.1, providing 
developers with endpoints, authentication methods, and code examples.
```

---

## 📋 Checklist

### Detected Type
`checklist`

### Typical Metadata
```json
{
  "title": "Pre-deployment Checklist",
  "date": "2024-01-15",
  "department": "DevOps",
  "items_count": "15",
  "status": "In Progress"
}
```

### Summary Example
```
Pre-deployment checklist with 15 items covering code review, testing, 
security scans, and infrastructure preparation.
```

---

## 🎯 How It Works

The AI analyzes the document and:

1. **Identifies Structure**: Looks for headers, formatting, key phrases
2. **Detects Type**: Classifies based on content and format patterns
3. **Extracts Metadata**: Finds dates, names, amounts, and other key data
4. **Generates Summary**: Creates a 2-3 sentence overview

### Flexible Detection

The system uses **GPT-4o-mini** (or your chosen model) which can:
- Adapt to various document formats
- Handle multiple languages
- Recognize industry-specific documents
- Extract custom fields based on content

### Customization

You can customize the extraction by modifying the prompt in:
```
external/thirdparty/openrouter/client.go
```

Look for the `promptTemplate` variable to adjust what information is extracted.

---

## 💡 Tips for Best Results

1. **Clear Documents**: Higher quality scans produce better results
2. **Standard Formats**: Common document types are recognized more accurately
3. **Structured Content**: Documents with clear sections work best
4. **Legible Text**: Ensure text is readable (avoid handwritten notes)

---

## 🔍 Testing Different Document Types

Upload various documents to test:

```bash
# Invoice
curl -X POST http://localhost:8080/documents/upload -F "file=@invoice.pdf"

# Resume
curl -X POST http://localhost:8080/documents/upload -F "file=@resume.docx"

# Report
curl -X POST http://localhost:8080/documents/upload -F "file=@report.pdf"
```

Each will be analyzed and classified accordingly!
