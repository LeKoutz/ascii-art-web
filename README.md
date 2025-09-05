# ASCII Art Web

A web application that converts text into ASCII art using different banner styles. Users can input text and select from various fonts to generate stylized ASCII art output.

## How to Run

### Local Development
```bash
go run main.go
```

## Docker

#### Quick Start
```bash
# Build and run with the provided script (Linux/Mac)
./docker-build.sh build
./docker-build.sh run
```

#### Audit commands
```bash
# Build the image
docker image build -t aaw-image .

# View images
docker images

# Run the container
docker container run -p 8081:8080 --detach --name aaw-container aaw-image

# View containers
docker ps -a

# Start an interactive shell section inside container
docker exec -it aaw-container /bin/sh

ls -l

# Inspect metadata
docker inspect aaw-image
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
├── docker-build.sh
├── Dockerfile
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
docker build -t aaw .

# Run container (detached mode)
docker run -d -p 8081:8080 --name aaw-container aaw

# Run container (interactive mode to see logs)
docker run -p 8081:8080 --name aaw-container aaw

# View logs
docker logs aaw-container
docker logs -f aaw-container  # Follow logs

# Stop container
docker stop aaw-container

# Remove container
docker rm aaw-container

# Remove image
docker rmi aaw

# Clean up unused Docker objects
docker system prune -f
```

### Docker Troubleshooting

**Port already in use:**
```bash
# Find what's using port 8080
netstat -tulpn | grep 8080
# Or use different port
docker run -p 8081:8080 --name aaw-container aaw
```

**Container name conflict:**
```bash
# Remove existing container
docker rm -f container
```

# **Contributors:**

- ### Constantine Ktistakis
- ### Giorgos Koutzos
- ### Ioannis Kountouris