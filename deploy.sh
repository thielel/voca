#!/bin/bash
set -e

# Voca Deployment Script for Google Cloud
# Usage: ./deploy.sh [backend|frontend|all]

PROJECT_ID="${GCP_PROJECT_ID:-voca-486110}"
REGION="${GCP_REGION:-europe-west1}"
BACKEND_SERVICE="voca-api"
OPENAI_API_KEY="${OPENAI_API_KEY:-}"

echo "🚀 Deploying Voca to Google Cloud"
echo "   Project: $PROJECT_ID"
echo "   Region: $REGION"
echo ""

deploy_backend() {
    echo "📦 Deploying Backend to Cloud Run..."
    cd backend

    # Check for OpenAI API key
    if [ -z "$OPENAI_API_KEY" ]; then
        echo "⚠️  Warning: OPENAI_API_KEY not set. AI interpretations will be disabled."
        echo "   Set it with: export OPENAI_API_KEY=sk-..."
    fi

    gcloud run deploy $BACKEND_SERVICE \
        --source . \
        --project $PROJECT_ID \
        --region $REGION \
        --allow-unauthenticated \
        --set-env-vars="ENVIRONMENT=production,DATABASE_PATH=/data/voca.db,OPENAI_API_KEY=$OPENAI_API_KEY" \
        --execution-environment gen2 \
        --cpu 1 \
        --memory 512Mi \
        --min-instances 0 \
        --max-instances 1 \
        --add-volume name=data-volume,type=cloud-storage,bucket=${PROJECT_ID}-data \
        --add-volume-mount volume=data-volume,mount-path=/data

    # Get the backend URL
    BACKEND_URL=$(gcloud run services describe $BACKEND_SERVICE --project $PROJECT_ID --region $REGION --format 'value(status.url)')
    echo ""
    echo "✅ Backend deployed: $BACKEND_URL"
    echo ""
    cd ..
}

deploy_frontend() {
    echo "📦 Building Frontend for Static Hosting..."
    cd frontend

    # Get backend URL if not provided
    if [ -z "$BACKEND_URL" ]; then
        BACKEND_URL=$(gcloud run services describe $BACKEND_SERVICE --project $PROJECT_ID --region $REGION --format 'value(status.url)' 2>/dev/null || echo "")
    fi

    if [ -z "$BACKEND_URL" ]; then
        echo "❌ Error: Backend URL not found. Deploy backend first or set BACKEND_URL."
        exit 1
    fi

    echo "   Using API URL: $BACKEND_URL"

    # Build with API URL
    NUXT_PUBLIC_API_URL=$BACKEND_URL npm run generate

    # Deploy to Firebase Hosting
    echo ""
    echo "🔥 Deploying to Firebase Hosting..."
    
    # Check if firebase is installed
    if ! command -v firebase &> /dev/null; then
        echo "Installing Firebase CLI..."
        npm install -g firebase-tools
    fi

    firebase deploy --only hosting --project $PROJECT_ID

    echo ""
    echo "✅ Frontend deployed to Firebase Hosting"
    cd ..
}

# Parse arguments
case "${1:-all}" in
    backend)
        deploy_backend
        ;;
    frontend)
        deploy_frontend
        ;;
    all)
        deploy_backend
        deploy_frontend
        ;;
    *)
        echo "Usage: $0 [backend|frontend|all]"
        exit 1
        ;;
esac

echo ""
echo "🎉 Deployment complete!"
