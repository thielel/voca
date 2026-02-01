# Deployment Guide - Google Cloud

This guide explains how to deploy Voca to Google Cloud using **Cloud Run** (backend) and **Firebase Hosting** (frontend).

## Cost Estimate

| Service | Monthly Cost |
|---------|-------------|
| Cloud Run (low traffic) | $0-2 |
| Firebase Hosting | Free |
| **Total** | **~$0-5/month** |

## Prerequisites

1. [Google Cloud account](https://cloud.google.com/)
2. [gcloud CLI](https://cloud.google.com/sdk/docs/install) installed
3. [Firebase CLI](https://firebase.google.com/docs/cli) installed (or install via npm)
4. Node.js 18+ and npm

## Setup

### 1. Create a Google Cloud Project

```bash
# Create a new project (or use existing)
gcloud projects create your-project-id --name="Voca"

# Set the project
gcloud config set project your-project-id

# Enable required APIs
gcloud services enable run.googleapis.com
gcloud services enable cloudbuild.googleapis.com
```

### 2. Configure Firebase

```bash
# Login to Firebase
firebase login

# Initialize Firebase in the frontend folder (select Hosting)
cd frontend
firebase init hosting

# When prompted:
# - Select "Use an existing project" → your-project-id
# - Public directory: .output/public
# - Single-page app: Yes
# - Automatic builds: No
```

Update `frontend/.firebaserc` with your project ID:

```json
{
  "projects": {
    "default": "your-project-id"
  }
}
```

### 3. Set Environment Variables

Create a `.env` file in the project root (not committed to git):

```bash
# Required for AI interpretations
OPENAI_API_KEY=sk-your-api-key-here

# Google Cloud settings
GCP_PROJECT_ID=your-project-id
GCP_REGION=europe-west1
```

## Deployment

### Option A: Use the Deploy Script

```bash
# Set your environment variables
export GCP_PROJECT_ID=your-project-id
export OPENAI_API_KEY=sk-your-key

# Deploy everything
./deploy.sh all

# Or deploy individually
./deploy.sh backend
./deploy.sh frontend
```

### Option B: Manual Deployment

#### Deploy Backend to Cloud Run

```bash
cd backend

gcloud run deploy voca-api \
  --source . \
  --project your-project-id \
  --region europe-west1 \
  --allow-unauthenticated \
  --set-env-vars="ENVIRONMENT=production,DATABASE_PATH=/data/voca.db,OPENAI_API_KEY=sk-xxx" \
  --execution-environment gen2 \
  --cpu 1 \
  --memory 512Mi \
  --min-instances 0 \
  --max-instances 3

# Note the URL output (e.g., https://voca-api-xxxxx.run.app)
```

#### Deploy Frontend to Firebase Hosting

```bash
cd frontend

# Build with the backend URL
NUXT_PUBLIC_API_URL=https://voca-api-xxxxx.run.app npm run generate

# Deploy to Firebase
firebase deploy --only hosting --project your-project-id
```

## Important Notes

### SQLite Persistence

Cloud Run containers are **stateless** by default. This project uses **Cloud Storage FUSE** for persistent SQLite storage.

**Current Setup:**
- Bucket: `gs://voca-486110-data`
- Mount path: `/data`
- Database file: `/data/voca.db`

The deploy script automatically mounts the Cloud Storage bucket. To set this up for a new project:

```bash
# 1. Create the storage bucket
gcloud storage buckets create gs://YOUR_PROJECT_ID-data --location=europe-west1

# 2. Deploy with volume mount (included in deploy.sh)
gcloud run deploy voca-api \
  --add-volume name=data-volume,type=cloud-storage,bucket=YOUR_PROJECT_ID-data \
  --add-volume-mount volume=data-volume,mount-path=/data
```

**Alternative: Cloud SQL** (for higher reliability)
- Migrate from SQLite to PostgreSQL
- Use Cloud SQL (~$7-10/month for smallest instance)

### Custom Domain

```bash
# For Firebase Hosting
firebase hosting:channel:deploy production

# For Cloud Run
gcloud run services update voca-api --domain your-domain.com
```

### Viewing Logs

```bash
# Backend logs
gcloud run logs read voca-api --region europe-west1

# Or in Cloud Console
# https://console.cloud.google.com/run
```

## Architecture

```
┌──────────────────────────────────┐
│      Firebase Hosting (CDN)       │
│        Static Frontend            │
│         (Nuxt SSG)                │
└──────────────┬───────────────────┘
               │ API calls
               ▼
┌──────────────────────────────────┐
│         Cloud Run                 │
│        Go Backend                 │
│     (scales to zero)              │
└──────────────┬───────────────────┘
               │
               ▼
┌──────────────────────────────────┐
│       SQLite Database             │
│   (or Cloud SQL for production)   │
└──────────────────────────────────┘
```
