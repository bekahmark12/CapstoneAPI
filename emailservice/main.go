package main

import (
	"log"
	"os"

	"github.com/yhung-mea7/SEN300-micro/tree/main/emailservice/messaging"
)

func main() {
	logger := log.New(os.Stdout, "cart-service", log.LstdFlags)

	consumer := messaging.NewRabbitMessanger(
		os.Getenv("RABBIT_CONN"),
		os.Getenv("SMTP_USER"),
		os.Getenv("SMTP_PASS"),
		os.Getenv("SMTP_HOST"),
		os.Getenv("SMTP_PORT"),
		logger,
	)

	consumer.Consume()
}
