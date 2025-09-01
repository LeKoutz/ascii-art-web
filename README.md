# ASCII Art Web

A web application that converts text into ASCII art using different banner styles. Users can input text and select from various fonts to generate stylized ASCII art output.

## How to Run

### Local Development
```bash
go run main.go
```

### Docker (Recommended)

#### Using Docker Compose
```bash
docker-compose up --build
```

#### Using Docker Commands
```bash
# Build the image
docker build -t ascii-art-web .

# Run the container
docker run -d -p 8080:8080 --name ascii-art-web-container ascii-art-web
```

#### Using Build Scripts
**Windows:**
```cmd
docker-build.bat build
docker-build.bat run
```

**Linux/Mac:**
```bash
chmod +x docker-build.sh
./docker-build.sh build
./docker-build.sh run
```

The server will start on `http://localhost:8080`

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
- **Non-root user** for security
- **Health checks** for container monitoring
- **Metadata labels** for better organization
- **Alpine Linux** base for minimal footprint

### Docker Commands

```bash
# Build image
docker build -t ascii-art-web .

# Run container
docker run -d -p 8080:8080 --name ascii-art-web-container ascii-art-web

# View logs
docker logs ascii-art-web-container

# Stop container
docker stop ascii-art-web-container

# Remove container
docker rm ascii-art-web-container

# Clean up unused Docker objects
docker system prune -f
```

## Contributors

- ### Constantine Ktistakis
- ### Giorgos Koutzos
- ### Ioannis Kountouris