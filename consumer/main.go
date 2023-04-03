package main

import (
	"fmt"
	"github.com/streadway/amqp"
	"log"
	"os"
)

func main() {
	amqURL := os.Getenv("AMQ_URL")
	if amqURL == "" {
		amqURL = "amqp://guest:guest@localhost:5672"
	}
	conn, err := amqp.Dial(amqURL)
	if err != nil {
		log.Fatalln("unable to connect amql", err)
	}

	channel, err := conn.Channel()
	if err != nil {
		log.Fatalln("unable to create channel")
	}
	_, err = channel.QueueDeclare("my_queue", true, false, false, false, nil)

	if err != nil {
		log.Fatalln("unable to declare queueu", err)
	}

	err = channel.QueueBind("my_queue", "#", "myChannel", false, nil)
	if err != nil {
		log.Fatalln("unable to bind queue", err)
	}
	message, err := channel.Consume("my_queue", "", false, false, false, false, nil)
	if err != nil {
		log.Fatalln("unable to consume message")
	}

	for msg := range message {
		fmt.Println("message received:", string(msg.Body))
		msg.Ack(false)
	}
	defer conn.Close()
}
