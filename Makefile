# Define the image names and container names
IMAGE_NAME = leader_app
DB_IMAGE_NAME = leader_db
CONTAINER_NAME = leader_app
DB_CONTAINER_NAME = leader_db

# Docker Compose command
DC = docker-compose

# Build the Docker images
build:
	$(DC) build

# Start the containers
up: build
	$(DC) up -d

# Stop and remove the containers
down:
	$(DC) down

# View logs for the application
logs:
	$(DC) logs -f

# Remove all containers and images
clean:
	$(DC) down --rmi all --volumes --remove-orphans

# Run tests (if any)
test:
	@echo "No tests defined in Makefile."

# Show help message
help:
	@echo "Makefile commands:"
	@echo "  build         - Build the Docker images"
	@echo "  up            - Build and start the Docker containers"
	@echo "  down          - Stop and remove the Docker containers"
	@echo "  logs          - View logs for the Docker containers"
	@echo "  clean         - Remove all Docker containers and images"
	@echo "  test          - Run tests (if any)"
	@echo "  help          - Show this help message"

.PHONY: build up down logs clean test help
