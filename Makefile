BINARY_NAME = schedule-manager-backend

build:
	go build -o ${BINARY_NAME} ./cmd/app

run: build
	docker-compose up --remove-orphans app

debug: build
	docker-compose up --remove-orphans debug