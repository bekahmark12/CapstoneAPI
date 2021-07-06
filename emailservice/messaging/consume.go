package messaging

import (
	"encoding/json"
	"log"
	"net/smtp"

	"github.com/streadway/amqp"
)

type (
	message struct {
		Name    string `json:"name"`
		Email   string `json:"email"`
		Content string `json:"content"`
	}

	emailConfig struct {
		from     string
		password string
		smtpHost string
		smtpPort string
	}

	RabbitMessanger struct {
		channel *amqp.Channel
		queue   *amqp.Queue
		email   *emailConfig
		log     *log.Logger
	}
)

func NewRabbitMessanger(rabbitConnStr string, from string, password string, smtpHost string, smtpPort string, l *log.Logger) *RabbitMessanger {
	conn, err := amqp.Dial(rabbitConnStr)
	if err != nil {
		panic(err)
	}
	ch, err := conn.Channel()
	if err != nil {
		panic(err)
	}
	q, err := ch.QueueDeclare(
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
	return &RabbitMessanger{
		channel: ch,
		queue:   &q,
		email:   &emailConfig{from, password, smtpHost, smtpPort},
		log:     l,
	}
}

func (r *RabbitMessanger) Consume() {
	defer r.channel.Close()
	msgs, err := r.channel.Consume(
		r.queue.Name,
		"",
		false,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		panic(err)
	}

	listen := make(chan bool)

	go func() {
		for d := range msgs {

			messageInfo := message{}
			if err := json.Unmarshal(d.Body, &messageInfo); err != nil {
				r.log.Fatal(err)
			}
			if err := r.sendEmail(&messageInfo); err != nil {
				r.log.Fatal(err)
			}
			d.Ack(false)
		}
	}()
	r.log.Println("Running consumer")
	<-listen
}

func (r *RabbitMessanger) sendEmail(message *message) error {
	auth := smtp.PlainAuth("", r.email.from, r.email.password, r.email.smtpHost)
	if err := smtp.SendMail(r.email.smtpHost+":"+r.email.smtpPort, auth, r.email.from, []string{message.Email}, []byte(message.Content)); err != nil {
		return err
	}
	return nil
}
