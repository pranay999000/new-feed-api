package services

import (
	"context"
	"log"
	"time"

	"github.com/rabbitmq/amqp091-go"
)

func FailOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s : %s", msg, err)
	}
}


func CreateUserNode(user_id string) {

	mq, err := amqp091.Dial("amqp://mq:password@localhost:5672/")
	FailOnError(err, "Failed to open channel")
	defer mq.Close()

	ch, err := mq.Channel()
	FailOnError(err, "Failed to open channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"hello",
		false,
		false,
		false,
		false,
		nil,
	)

	FailOnError(err, "Failed to declare a queue")

	ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
	defer cancel()

	err = ch.PublishWithContext(ctx,
		"",
		q.Name,
		false,
		false,
		amqp091.Publishing{
			ContentType: "text/plain",
			Body: []byte(user_id),
		},
	)

	FailOnError(err, "Failed to publish a message")
}

