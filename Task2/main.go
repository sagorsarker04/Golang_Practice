package main

import (
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	// Step 1: Connect to RabbitMQ
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		log.Fatal("❌ Failed to connect to RabbitMQ:", err)
	}
	defer conn.Close()

	// Step 2: Open a channel
	ch, err := conn.Channel()
	if err != nil {
		log.Fatal("❌ Failed to open a channel:", err)
	}
	defer ch.Close()

	// Step 3: Declare the same queue your consumer is using
	queueName := "email-sender-queue" // 🔁 Must match the consumer queue name
	_, err = ch.QueueDeclare(
		queueName,
		true, // durable
		false, // auto-delete
		false, // exclusive
		false, // no-wait
		nil,   // args
	)
	if err != nil {
		log.Fatal("❌ Failed to declare queue:", err)
	}

	// Step 4: Prepare your message (your provided structure)
	body := `
{
  "to": ["ramita.rafi@gmail.com"],
  "subject": "Your Subject",
  "template": "generic_email.html",
  "data": {
    "Body": "This is the dynamic content of the email."
  }
}`

	// Step 5: Publish the message
	err = ch.Publish(
		"",        // exchange: default
		queueName, // routing key = queue name
		false, false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        []byte(body),
		},
	)
	if err != nil {
		log.Fatal("❌ Failed to publish message:", err)
	}

	log.Println("✅ Email event sent successfully.")
}
