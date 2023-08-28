BINARY_NAME='event_delivery'
IMAGE_NAME = 'event-delivery-app'

build: build-go-server

build-go-server:
	rm -f $(BINARY_NAME)
	go build -o $(BINARY_NAME) cmd/main.go

start:
	go run cmd/main.go start

docker-build:
	docker build -t $(IMAGE_NAME) .

docker-up:
	docker-compose up

deploy: docker-build docker-up