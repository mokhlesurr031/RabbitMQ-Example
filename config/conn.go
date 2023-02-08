package config

import (
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

func GetMQConnection(taskName string) (ch *amqp.Channel, queue amqp.Queue, err error) {
	mq := MQ()
	conn, err := amqp.Dial(mq.Broker)
	failOnError(err, "Failed to connect to RabbitMQ")
	//defer conn.Close()

	ch, err = conn.Channel()
	failOnError(err, "Failed to open a channel")
	//defer ch.Close()

	q, err := ch.QueueDeclare(
		taskName, // name
		false,    // durable
		false,    // delete when unused
		false,    // exclusive
		false,    // no-wait
		nil,      // arguments
	)
	failOnError(err, "Failed to declare a queue")

	return ch, q, nil
}
