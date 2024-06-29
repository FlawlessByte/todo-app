build:
	go mod download && CGO_ENABLED=0 GOOS=linux go build -o ./cmd/main.go

run: build
	docker-compose up --build todo-app