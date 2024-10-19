.PHONY: all start stop product db

run:
	docker-compose up --build

run-dev:
	docker-compose -f docker-compose.dev.yaml up --build

# Stop all services and databases
stop:
	docker-compose down
	pkill -f 'air'

# gen keyaprs
gen-key:
	openssl genrsa -out rsa_private.pem 2048
	openssl rsa -in rsa_private.pem -pubout -out rsa_public.pem
