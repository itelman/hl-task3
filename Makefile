# Define variables for docker-compose
DOCKER_COMPOSE = docker-compose

# Default target to build and run the services
.PHONY: up
up:
	$(DOCKER_COMPOSE) up --build -d

# Target to stop and remove containers, networks, volumes, and images
.PHONY: down
down:
	$(DOCKER_COMPOSE) down

# Target to view logs from all services
.PHONY: logs
logs:
	$(DOCKER_COMPOSE) logs -f

# Target to access the Go app container's shell
.PHONY: app-sh
app-sh:
	$(DOCKER_COMPOSE) exec app /bin/sh

# Target to access the PostgreSQL container's shell
.PHONY: db-sh
db-sh:
	$(DOCKER_COMPOSE) exec db /bin/bash
