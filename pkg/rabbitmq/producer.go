package rabbitmq

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/streadway/amqp"
)

type RabbitMQProducer struct {
	RabbitMQClient
}

func NewRabbitMQProducer(client *RabbitMQClient) *RabbitMQProducer {
	p := &RabbitMQProducer{
		RabbitMQClient: *client,
	}
	return p
}

func (r *RabbitMQProducer) Publish(file string) {

	data, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Errorf(err.Error())
	}
	var jsonArr []interface{}
	err = json.Unmarshal(data, &jsonArr)
	if err != nil {
		fmt.Errorf(err.Error())
	}
	for _, obj := range jsonArr {
		data, _ := json.Marshal(obj)
		err := r.ch.Publish(
			r.exchangeName, // exchange
			r.q.Name,       // routing key
			false,          // mandatory
			false,          // immediate
			amqp.Publishing{
				DeliveryMode: amqp.Persistent,
				ContentType:  "text/plain",
				Body:         []byte(data),
			})
		if err != nil {
			fmt.Errorf(err.Error())
		}
	}

}
