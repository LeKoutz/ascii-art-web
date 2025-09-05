#!/bin/bash

# ASCII Art Web Docker Build Script
# This script builds the Docker image and provides various operations

set -e

IMAGE_NAME="aaw-image"
CONTAINER_NAME="aaw-container"
TAG="latest"

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

print_usage() {
    echo "Usage: $0 [build|run|stop|clean|logs|shell]"
    echo ""
    echo "Commands:"
    echo "  build  - Build the Docker image" 
    echo "  run    - Run the container"
    echo "  stop   - Stop the container"
    echo "  clean  - Remove container and image"
    echo "  logs   - Show container logs"
    echo "  shell  - Open shell in running container"
}

build_image() {
    echo -e "${GREEN}Building Docker image...${NC}"
    docker build -t ${IMAGE_NAME}:${TAG} .
    echo -e "${GREEN}Image built successfully!${NC}"
    
    # Show image details
    echo -e "${YELLOW}Image details:${NC}"
    docker images ${IMAGE_NAME}:${TAG}
}

run_container() {
    echo -e "${GREEN}Running container...${NC}"
    
    # Stop existing container if running
    if docker ps -q -f name=${CONTAINER_NAME} | grep -q .; then
        echo -e "${YELLOW}Stopping existing container...${NC}"
        docker stop ${CONTAINER_NAME}
        docker rm ${CONTAINER_NAME}
    fi
    
    # Run new container
    docker run -d \
        --name ${CONTAINER_NAME} \
        -p 8080:8080 \
        --restart unless-stopped \
        ${IMAGE_NAME}:${TAG}
    
    echo -e "${GREEN}Container started successfully!${NC}"
    echo -e "${YELLOW}Access the application at: http://localhost:8080${NC}"
    
    # Show container status
    docker ps -f name=${CONTAINER_NAME}
}

stop_container() {
    echo -e "${YELLOW}Stopping container...${NC}"
    docker stop ${CONTAINER_NAME} 2>/dev/null || echo "Container not running"
    docker rm ${CONTAINER_NAME} 2>/dev/null || echo "Container not found"
    echo -e "${GREEN}Container stopped and removed${NC}"
}

clean_all() {
    echo -e "${YELLOW}Cleaning up...${NC}"
    
    # Stop and remove container
    docker stop ${CONTAINER_NAME} 2>/dev/null || true
    docker rm ${CONTAINER_NAME} 2>/dev/null || true
    
    # Remove image
    docker rmi ${IMAGE_NAME}:${TAG} 2>/dev/null || true
    
    # Clean up unused images and containers
    docker system prune -f
    
    echo -e "${GREEN}Cleanup completed${NC}"
}

show_logs() {
    echo -e "${GREEN}Container logs:${NC}"
    docker logs -f ${CONTAINER_NAME}
}

open_shell() {
    echo -e "${GREEN}Opening shell in container...${NC}"
    docker exec -it ${CONTAINER_NAME} /bin/sh
}

# Main script logic
case "${1:-}" in
    build)
        build_image
        ;;
    run)
        run_container
        ;;
    stop)
        stop_container
        ;;
    clean)
        clean_all
        ;;
    logs)
        show_logs
        ;;
    shell)
        open_shell
        ;;
    *)
        print_usage
        exit 1
        ;;
esac
