# Heroku Deployment Guide

This guide will help you deploy the AI Document Summarizer service to Heroku.

## Prerequisites

1. **Heroku Account**: Sign up at [heroku.com](https://signup.heroku.com/)
2. **Heroku CLI**: Install from [devcenter.heroku.com/articles/heroku-cli](https://devcenter.heroku.com/articles/heroku-cli)
3. **Git Repository**: Your project must be in a git repository
4. **External S3 Storage**: Heroku's ephemeral filesystem requires external storage (AWS S3, Cloudflare R2, or similar)

## Required Services

### 1. S3-Compatible Storage

You need an S3-compatible service for document storage. Options include:

- **AWS S3**: [aws.amazon.com/s3](https://aws.amazon.com/s3/)
- **Cloudflare R2**: [developers.cloudflare.com/r2](https://developers.cloudflare.com/r2/) (Free tier available)
- **Backblaze B2**: [backblaze.com/b2](https://www.backblaze.com/b2/cloud-storage.html)
- **DigitalOcean Spaces**: [digitalocean.com/products/spaces](https://www.digitalocean.com/products/spaces)

Create a bucket and obtain:
- Access Key ID
- Secret Access Key
- Bucket name
- Endpoint URL (if not AWS S3)

### 2. OpenRouter API Key

Get your API key from [openrouter.ai](https://openrouter.ai/)

## Quick Deployment (Automated)

Use the Makefile for automated setup:

```bash
# Complete setup in one command
make heroku-setup

# Then deploy
make heroku-deploy
```

This will:
1. Log you into Heroku
2. Create a new app with container stack
3. Add PostgreSQL database
4. Set all required environment variables

## Manual Deployment (Step-by-Step)

### Step 1: Login to Heroku

```bash
heroku login
```

### Step 2: Create Heroku App

```bash
# Create with random name
heroku create --stack container

# Or create with specific name
heroku create your-app-name --stack container
```

### Step 3: Add PostgreSQL Database

```bash
heroku addons:create heroku-postgresql:essential-0
```

The database URL will be automatically set as `DATABASE_URL` environment variable. The essential-0 plan costs ~$5/month.

### Step 4: Set Environment Variables

```bash
# Required: OpenRouter API Key
heroku config:set OPENROUTER_API_KEY=your_openrouter_api_key

# Required: S3 Credentials
heroku config:set AWS_ACCESS_KEY_ID=your_aws_access_key
heroku config:set AWS_SECRET_ACCESS_KEY=your_aws_secret_key
heroku config:set S3_BUCKET=your_bucket_name

# Optional: S3 Endpoint (required for non-AWS S3)
heroku config:set S3_ENDPOINT=https://your-s3-endpoint.com

# Optional: Model and settings (with defaults)
heroku config:set OPENROUTER_MODEL=openai/gpt-4o-mini
heroku config:set AWS_REGION=us-east-1
heroku config:set MAX_FILE_SIZE_MB=5
```

### Step 5: Update Database Configuration

The app automatically detects Heroku's `DATABASE_URL`. No code changes needed.

### Step 6: Deploy to Heroku

```bash
# Make sure all changes are committed
git add .
git commit -m "Prepare for Heroku deployment"

# Push to Heroku
git push heroku main
```

Heroku will:
- Build the Docker container
- Install all dependencies (including PDF processing tools)
- Deploy the application

### Step 7: Scale the App

```bash
# Ensure at least one web dyno is running
heroku ps:scale web=1
```

### Step 8: Open Your App

```bash
heroku open
```

Or use the Makefile:
```bash
make heroku-open
```

## Verify Deployment

### Check App Status

```bash
heroku ps
```

### View Logs

```bash
# Using Heroku CLI
heroku logs --tail

# Or using Makefile
make heroku-logs
```

### Test the Health Endpoint

```bash
# Replace your-app-name with your actual app name
curl https://your-app-name.herokuapp.com/health
```

Expected response:
```json
{"status":"ok"}
```

### Test Document Upload

```bash
curl -X POST https://your-app-name.herokuapp.com/documents/upload \
  -F "file=@/path/to/your/document.pdf"
```

## Configuration Management

### View All Config Variables

```bash
heroku config

# Or using Makefile
make heroku-info
```

### Update Environment Variable

```bash
heroku config:set VARIABLE_NAME=new_value
```

### Remove Environment Variable

```bash
heroku config:unset VARIABLE_NAME
```

## Database Management

### Access PostgreSQL Database

```bash
heroku pg:psql
```

### View Database Info

```bash
heroku pg:info
```

### Database Backups

```bash
# Create manual backup
heroku pg:backups:capture

# View backups
heroku pg:backups
```

## Scaling

### Vertical Scaling (Dyno Size)

```bash
# Upgrade to Standard-1X
heroku ps:type web=standard-1x

# Upgrade to Standard-2X (more memory/CPU)
heroku ps:type web=standard-2x
```

### Horizontal Scaling (Multiple Dynos)

```bash
# Scale to 2 dynos
heroku ps:scale web=2
```

## Monitoring

### View App Metrics

```bash
heroku apps:info
```

### View Recent Logs

```bash
heroku logs --num 100
```

### View Specific Component Logs

```bash
heroku logs --source app --tail
```

## Troubleshooting

### App Not Starting

1. Check logs:
   ```bash
   heroku logs --tail
   ```

2. Verify environment variables:
   ```bash
   heroku config
   ```

3. Ensure all required variables are set:
   - `OPENROUTER_API_KEY`
   - `AWS_ACCESS_KEY_ID`
   - `AWS_SECRET_ACCESS_KEY`
   - `S3_BUCKET`

### Database Connection Issues

1. Verify PostgreSQL addon is installed:
   ```bash
   heroku addons
   ```

2. Check database URL:
   ```bash
   heroku config:get DATABASE_URL
   ```

### File Upload Issues

1. Verify S3 credentials are correct
2. Check bucket permissions
3. For non-AWS S3, ensure `S3_ENDPOINT` is set correctly

### Build Failures

1. Check Dockerfile syntax
2. Verify go.mod dependencies are up to date
3. Review build logs:
   ```bash
   heroku builds:info
   ```

## Cost Optimization

### Free Tier Limitations

- **Dynos**: Apps sleep after 30 minutes of inactivity
- **PostgreSQL**: essential-0 addon (~$5/month) - 10,000 rows
- **Build Time**: Limited build minutes per month

### Recommendations

1. **Use Hobby Dynos ($7/month)**: Prevents sleeping
2. **Optimize Docker Image**: Remove unnecessary files (see `.slugignore`)
3. **Monitor Usage**: Use Heroku metrics dashboard

## Updating the App

### Deploy Updates

```bash
# Commit your changes
git add .
git commit -m "Your update message"

# Deploy
git push heroku main
```

### Rollback to Previous Version

```bash
# View releases
heroku releases

# Rollback to previous version
heroku rollback
```

## Custom Domain

### Add Custom Domain

```bash
heroku domains:add www.yourdomain.com
```

### Configure DNS

Point your domain to the Heroku DNS target shown after adding the domain.

### Enable SSL (Automatic)

Heroku provides free automated SSL certificates for custom domains.

## Makefile Commands Reference

All Heroku operations have Makefile shortcuts:

```bash
make heroku-login        # Login to Heroku
make heroku-create       # Create new app
make heroku-addons       # Add PostgreSQL
make heroku-config       # Set environment variables
make heroku-deploy       # Deploy to Heroku
make heroku-logs         # View logs (streaming)
make heroku-open         # Open app in browser
make heroku-setup        # Complete setup (all above)
make heroku-info         # Show app and config info
```

## Security Best Practices

1. **Never commit API keys**: Use environment variables only
2. **Use strong S3 bucket policies**: Restrict public access
3. **Enable CORS carefully**: Only allow trusted origins
4. **Monitor usage**: Set up billing alerts
5. **Regular backups**: Schedule database backups
6. **Review logs**: Check for suspicious activity

## Support

- **Heroku Documentation**: [devcenter.heroku.com](https://devcenter.heroku.com/)
- **Heroku Status**: [status.heroku.com](https://status.heroku.com/)
- **OpenRouter Support**: [openrouter.ai](https://openrouter.ai/)

## Next Steps

After deployment:

1. Test all three endpoints (upload, analyze, retrieve)
2. Monitor logs for any errors
3. Set up monitoring/alerting
4. Configure custom domain (optional)
5. Enable backups
6. Document your API URL for client applications
