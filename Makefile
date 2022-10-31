include .env

.SILENT:

build:
	docker-compose -f docker-compose.yml build csvloader

run:
	docker-compose -f docker-compose.yml up -d csvloader

stop:
	docker-compose -f docker-compose.yml down

migrate:
	docker-compose -f docker-compose.yml run --rm csvloader migrate -path ./schema -database ${DATABASE_URL} up

migrate-down:
	docker-compose -f docker-compose.yml run --rm csvloader migrate -path ./schema -database ${DATABASE_URL} down
