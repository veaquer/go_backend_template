# Stage 1: Base image with Air & dependencies
FROM golang:1.24

# Install necessary packages and Air
RUN apt-get update && apt-get install -y git curl && \
    go install github.com/air-verse/air@v1.62.0

# Create app directory
WORKDIR /app

# Copy go.mod and go.sum early to cache deps
COPY go.mod go.sum ./
RUN go mod download

# Copy all sources (excluding via .dockerignore if needed)
COPY . .

# Ensure tmp dir for Air exists (matches .air.toml)
RUN mkdir -p tmp

# Expose app port (matches .env)
EXPOSE 4000

# Command to run Air
CMD ["air"]
