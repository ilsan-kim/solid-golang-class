package rabbitmq

import (
	"log"

	"github.com/streadway/amqp"
)

type RabbitMQClient struct {
	exchangeName string
	conn         *amqp.Connection
	ch           *amqp.Channel
	q            amqp.Queue
}

func NewRabbitMQClient(host string) *RabbitMQClient {
	conn, err := amqp.Dial(host)
	failOnError(err, "Failed to connect to RabbitMQ")

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")

	rc := &RabbitMQClient{
		conn: conn,
		ch:   ch,
	}
	return rc
}
func (r *RabbitMQClient) Close() {
	r.conn.Close()
	r.ch.Close()
}
func (r *RabbitMQClient) QueueDeclare(name string) {
	q, err := r.ch.QueueDeclare(
		name,  // name
		true,  // durable
		false, // delete when unused
		false, // exclusive
		false, // no-wait
		nil,   // arguments
	)
	failOnError(err, "Failed to declare a queue")
	r.q = q
}

func (r *RabbitMQClient) QueueBind() {
	err := r.ch.QueueBind(
		r.q.Name,       // queue name
		r.q.Name,       // routing key
		r.exchangeName, // exchange
		false,
		nil,
	)
	failOnError(err, "Failed to bind a queue")
}

func (r *RabbitMQClient) ExchangeDeclare(name string, kind string) {
	r.exchangeName = name
	err := r.ch.ExchangeDeclare(
		name,  // name
		kind,  // type
		true,  // durable
		false, // auto-deleted
		false, // internal
		false, // no-wait
		nil,   // arguments
	)
	failOnError(err, "Failed to declare an exchange")
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}
