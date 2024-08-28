
setup: docker-compose.yaml Dockerfile go.mod go.sum .env
	docker-compose up --build

swagger:
	swag init -d cmd/chater-backend/,internal/api,internal/domain/models

build: go.mod go.sum
	CGO_ENABLED=0 GOOS=linux go build -o ./ cmd/chater-backend/main.go
