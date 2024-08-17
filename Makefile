.PHONY: all start stop product db

# Start all services and databases
all: db start

# Start services
start: product auth

# Start product with hot reload
product:
	cd services/product && air

# Start auth with hot reload
auth:
	cd services/auth && air

# Start databases using Docker Compose
db:
	docker-compose up -d

# Stop all services and databases
stop:
	docker-compose down
	pkill -f 'air'
