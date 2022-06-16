package main

import (
	"flag"
	"mock-data-generator/pkg/rabbitmq"
)

func main() {
	host := flag.String("host", "amqp://guest:guest@:::5672/", "rabbitmq host you are connecting to")
	queue := flag.String("queue", "users", "queue name you are publishing data to")
	filePath := flag.String("file path of mock json data", "mock_data.json", "data file you are loading to publish")

	flag.Parse()
	client := rabbitmq.NewRabbitMQClient(*host)
	p := rabbitmq.NewRabbitMQProducer(client)
	p.QueueDeclare(*queue)
	p.Publish(*filePath)
	defer p.Close()
}
