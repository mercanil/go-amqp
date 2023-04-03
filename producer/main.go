package main

import (
	"github.com/streadway/amqp"
	"log"
	"os"
)

func main() {
	amqpUrl := os.Getenv("AMQP_URL")
	if amqpUrl == "" {
		amqpUrl = "amqp://guest:guest@localhost:5672"

	}
	connection, err := amqp.Dial(amqpUrl)
	if err != nil {
		log.Fatalln("Error on connection")
	}
	defer connection.Close()

	channel, err := connection.Channel()
	if err != nil {
		log.Fatalln("error creating channel")

	}
	err = channel.ExchangeDeclare("myChannel", "topic", true, false, false, false, nil)

	if err != nil {
		log.Fatalln("unable to declare exchange" + err.Error())
	}

	message := amqp.Publishing{
		Body: []byte("Hello World"),
	}
	err = channel.Publish("myChannel", "aNewMessageKey", false, false, message)

	if err != nil {
		log.Fatalln("unable to send message" + err.Error())

	}

}
