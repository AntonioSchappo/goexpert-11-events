package main

import (
	"context"

	"github.com/AntonioSchappo/eventlib/pkg/rabbitmq"
)

func main() {
	ch, err := rabbitmq.OpenChannel()
	if err != nil {
		panic(err)
	}
	defer ch.Close()

	ctx := context.Background()

	rabbitmq.Publish(ch, ctx, "Hello World!", "amq.direct")
}
