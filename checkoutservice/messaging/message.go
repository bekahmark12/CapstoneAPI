package messaging

import (
	"encoding/json"
	"fmt"

	"github.com/streadway/amqp"
)

type Messager struct {
	channel *amqp.Channel
	queue   *amqp.Queue
}

type Message struct {
	Name    string `json:"name"`
	Email   string `json:"email"`
	Content string `json:"content"`
}

func NewMessager(dialInfo string) *Messager {
	conn, err := amqp.Dial(dialInfo)
	if err != nil {
		panic(err)
	}
	channel, err := conn.Channel()
	if err != nil {
		panic(err)
	}

	q, err := channel.QueueDeclare(
		"EmailQueue",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		panic(err)
	}

	return &Messager{
		channel: channel,
		queue:   &q,
	}
}

func (m *Messager) SubmitToMessageBroker(message *Message) error {
	body, err := json.Marshal(message)
	if err != nil {
		return err
	}

	if err := m.channel.Publish(
		"",
		m.queue.Name,
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		},
	); err != nil {
		fmt.Println(err.Error())
		return err
	}
	return nil
}
