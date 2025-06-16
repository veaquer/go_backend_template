#!/bin/bash

echo "=== Cleaning up existing containers and images ==="
docker-compose -f docker-compose.prod.yml down

echo "=== Removing old backend image to force rebuild ==="
docker rmi go_backend_template-backend 2>/dev/null || echo "No existing backend image found"

echo "=== Building production image without cache ==="
docker-compose -f docker-compose.prod.yml build --no-cache

echo "=== Starting production environment ==="
docker-compose --env-file .env.production -f docker-compose.prod.yml up -d 
