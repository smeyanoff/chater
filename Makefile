
.PHONY: .swag
.swag:
	cd ./backend && swag init

setup: docker-compose.yaml .swag
	docker-compose up --build -d
