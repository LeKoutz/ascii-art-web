# ASCII Art Web

A web application that converts text into ASCII art using different banner styles. Users can input text and select from various fonts to generate stylized ASCII art output.

## How to Run

### Local Development
```bash
go run main.go
```

### Docker

#### Quick Start
```bash
# Build and run with the provided script (Linux/Mac)
./docker-build.sh build
./docker-build.sh run
```

#### Manual Docker Commands
```bash
# Build the image
docker build -t ascii-art-web .

# Run the container
docker run -d -p 8080:8080 --name ascii-art-web-container ascii-art-web

# View logs
docker logs ascii-art-web-container

# Stop and remove
docker stop ascii-art-web-container
docker rm ascii-art-web-container
```

## Project Structure

```
ascii-art-web/
├── main.go              # Entry point
├── main_test.go         # Integration and main package tests
├── go.mod               # Go module definition
├── LICENSE              # MIT License
├── handlers/
│   ├── handlers.go      # HTTP request handlers
│   └── handlers_test.go # Handler unit tests
├── services/
│   ├── ascii-art.go     # ASCII art generation logic
│   └── ascii-art_test.go # Service unit tests
├── templates/
│   ├── index.html       # Main page template
│   └── error.html       # Error page template
└── banners/
    ├── standard.txt     # ASCII art fonts
    ├── shadow.txt
    ├── thinkertoy.txt
    └── zigzag.txt
```

## Error Status Testing

### 400 Bad Request
- Submit empty text
- Submit text with non-ASCII characters

### 404 Not Found
- Visit any URL other than `/` or `/ascii-art`
- Example: `http://localhost:8080/nonexistent`

### 500 Internal Server Error
- Delete or corrupt template files in `/templates/` directory

## Docker

This application is containerized using Docker with the following features:

- **Multi-stage build** for optimized image size
- **Alpine Linux** base for minimal footprint (~15MB final image)
- **Metadata labels** for better organization
- **Port 8080** exposed for web access

### Docker Build Script (Recommended)

The `docker-build.sh` script provides easy Docker management:

```bash
# Available commands:
./docker-build.sh build   # Build the Docker image
./docker-build.sh run     # Run the container (stops existing if running)
./docker-build.sh stop    # Stop and remove container
./docker-build.sh logs    # View container logs
./docker-build.sh shell   # Open shell in running container
./docker-build.sh clean   # Remove container and image completely
```

### Manual Docker Commands

```bash
# Build image
docker build -t ascii-art-web .

# Run container (detached mode)
docker run -d -p 8080:8080 --name ascii-art-web-container ascii-art-web

# Run container (interactive mode to see logs)
docker run -p 8080:8080 --name ascii-art-web-container ascii-art-web

# View logs
docker logs ascii-art-web-container
docker logs -f ascii-art-web-container  # Follow logs

# Stop container
docker stop ascii-art-web-container

# Remove container
docker rm ascii-art-web-container

# Remove image
docker rmi ascii-art-web

# Clean up unused Docker objects
docker system prune -f
```

### Docker Troubleshooting

**Port already in use:**
```bash
# Find what's using port 8080
netstat -tulpn | grep 8080
# Or use different port
docker run -p 8081:8080 --name ascii-art-web-container ascii-art-web
```

**Container name conflict:**
```bash
# Remove existing container
docker rm -f asci
```

# **Contributors:**

- ### Constantine Ktistakis
- ### Giorgos Koutzos
- ### Ioannis Kountouris