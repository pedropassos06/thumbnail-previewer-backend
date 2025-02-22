# Use official Go image as a build stage
FROM golang:1.22-alpine AS builder

# Set the working directory
WORKDIR /app

# Copy Go modules files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the application code
COPY . .

# Build the application
RUN go build -o thumbnail-previewer-backend ./cmd

# Final container stage
FROM alpine:latest

# Install dependencies for your app (e.g., bash, and any other required packages for dotenv)
RUN apk add --no-cache bash

# Set the working directory
WORKDIR /root/

# Copy built binary from builder
COPY --from=builder /app/thumbnail-previewer-backend .

# Copy the .env file to the container
COPY .env .env

# Command to run the application (ensure environment variables are loaded)
CMD ["/bin/sh", "-c", "source .env && ./thumbnail-previewer-backend"]