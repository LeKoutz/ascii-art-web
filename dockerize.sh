#!/bin/bash

# Build the Docker image
echo -e "Building Docker image..."
docker image build -f Dockerfile -t ascii-art-web-docker .

# Show images
echo -e "\n Current Docker images:"
docker images

# Run the container
echo -e "\n Running Docker container..."
docker container run -p 8080:8080 --detach --name dockerize ascii-art-web-docker

# Show containers
echo -e "\n Current Docker containers:"
docker ps -a

# Clean up dangling images
echo -e "\n Cleaning up..."
docker image prune -a -f

echo "Done! Container is running on http://localhost:8080"
echo "          To stop the container, use:                 docker container stop dockerize"
echo "          To remove the container, use:               docker container rm dockerize"
echo "          To clean up everything afterwards, use:     docker system prune -a --volumes"