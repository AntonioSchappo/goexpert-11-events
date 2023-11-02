# goexpert-11-events
A simple rabbitmq consumer and publisher implementation in go.

Instructions:

- Run command docker-compose up
- Set queues and exchanges accordingly at http://localhost:15672/
- Use command go run main.go in ./cmd/consumer in order to start the consumer
- Use command go run main.go in ./cmd/producer to send messages to the producer
