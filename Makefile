
setup: docker-compose.yaml
	docker-compose up --build

run_backend:
	go run ./backend/main.go
