package mq

import (
	"github.com/aryahadii/runandeh/configuration"
	"github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
)

var (
	conn    *amqp.Connection
	channel *amqp.Channel

	runsQueue amqp.Queue

	runsDelivery <-chan amqp.Delivery
)

// SubscribeToRunsQueue sets callback on deliveries
func SubscribeToRunsQueue(callback func(amqp.Delivery)) {
	go func() {
		for run := range runsDelivery {
			go callback(run)
		}
	}()
}

// InitMessageQueue creates connection, channel and queue
func InitMessageQueue() {
	// Make Connection
	var err error
	conn, err = amqp.Dial(configuration.GetInstance().GetString("rabbit-mq.url"))
	if err != nil {
		logrus.WithError(err).Fatal("can't connect to message queue")
	}

	// Make Channel
	channel, err = conn.Channel()
	if err != nil {
		logrus.WithError(err).Fatal("can't create message queue channel")
	}

	// Make Queue
	runsQueue, err = channel.QueueDeclare(
		configuration.GetInstance().GetString("rabbit-mq.runs-queue"), // name
		false, // durable
		false, // delete when unused
		false, // exclusive
		false, // no-wait
		nil,   // arguments
	)
	if err != nil {
		logrus.WithError(err).Fatal("can't create queue")
	}

	// Consumer
	runsDelivery, err = channel.Consume(
		runsQueue.Name, // queue
		"",             // consumer
		true,           // auto-ack
		false,          // exclusive
		false,          // no-local
		false,          // no-wait
		nil,            // args
	)
	if err != nil {
		logrus.WithError(err).Fatal("can't init consumer")
	}

	logrus.Info("message queue initialized")
}

// Close closes connection to RabbitMQ
func Close() {
	conn.Close()
}
