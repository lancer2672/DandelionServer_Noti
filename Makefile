rabbitmq:
	docker run -d --name dandelion   -p 5672:5672   -p 15672:15672   rabbitmq:latest
server:
	go run main.go
consumer: 
	go run cmd/consumer/main.go
publisher: 
	go run cmd/publisher/main.go
.PHONY: rabbitmq server consumer publisher