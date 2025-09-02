# Build stage
FROM golang:1.24.6-alpine AS builder

# Add metadata labels
LABEL maintainer="ASCII Art Web Team"
LABEL description="ASCII Art Web Application"
LABEL version="1.0.0"
LABEL org.opencontainers.image.source="https://platform.zone01.gr/git/cktistak/ascii-art-web/src/branch/dockerize/"
LABEL org.opencontainers.image.description="A web application that converts text into ASCII art"
LABEL org.opencontainers.image.licenses="MIT"

# Set working directory
WORKDIR /app

# Copy go mod files first for better caching
COPY go.mod go.sum* ./

# Copy source code
COPY . .

RUN go build -o /app/main 

# Final stage - minimal runtime image
FROM alpine:latest

# Add metadata to final image
LABEL maintainer="ASCII Art Web Team"
LABEL description="ASCII Art Web Application - Runtime"
LABEL version="1.0.0"

# Set working directory
WORKDIR /app

# Copy binary from builder stage
COPY --from=builder /app/main .

# Copy static files (templates and banners)
COPY --from=builder /app/templates ./templates
COPY --from=builder /app/banners ./banners

# Expose port
EXPOSE 8080

# Run the script
ENTRYPOINT ["/app/main"]
