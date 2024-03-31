package main

import (
	"fmt"

	amqp "github.com/rabbitmq/amqp091-go"

	"github.com/KurobaneShin/go-events/pkg/rabbitmq"
)

func main() {
	channel, err := rabbitmq.OpenChannel()
	if err != nil {
		panic(err)
	}

	defer channel.Close()

	msgs := make(chan amqp.Delivery)

	go rabbitmq.Consume(channel, msgs)

	for msg := range msgs {
		fmt.Println(string(msg.Body))
		err = msg.Ack(false)
	}
}
