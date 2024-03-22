rabbitmq:
	docker run -d --name dandelion   -p 5672:5672   -p 15672:15672   rabbitmq:latest

.PHONY: rabbitmq