package main

import "github.com/KurobaneShin/go-events/pkg/rabbitmq"

func main() {
	channel, err := rabbitmq.OpenChannel()
	if err != nil {
		panic(err)
	}

	defer channel.Close()

	rabbitmq.Publish(channel, "Hello, World!")
}
