package infraestructure

import (
	"api_event_driven_2/src/config"
	"encoding/json"
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

type Rabbit struct {
conn *amqp.Connection
}

func NewRabbitMQPublisher() *Rabbit {
    conn := config.GetRabbitMQConnection()
    return &Rabbit{conn}
}

func (r *Rabbit) SendConfirmation(message string) (*string,error) {
	ch, err := r.conn.Channel()
    if err != nil {
        return nil,err
    }
    defer ch.Close()

    err = ch.ExchangeDeclare(
        "courses", // name
        "direct", // type
        true,     // durable
        false,     // auto-deleted
        false,     // internal
        false,     // no-wait
        nil,       // arguments
    )
    if err != nil {
        return nil,err
    }

	body, err := json.Marshal(message)
	if err != nil {
        return nil, err
    }
    err = ch.Publish(
        "courses", // exchange
        "confirmation_key", // routing key
        false,  // mandatory
        false,  // immediate
        amqp.Publishing{
            ContentType: "text/plain",
            Body:        body,
        })
		if err != nil {
			return nil, err
		}
	
		log.Printf("[x] Sent %s", body)
		return &message, nil
}