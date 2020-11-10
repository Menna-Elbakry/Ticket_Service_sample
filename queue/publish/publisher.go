package publish

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"

	uuid "github.com/satori/go.uuid"
	"github.com/streadway/amqp"
)

func FailOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func RandomString(l int) string {
	bytes := make([]byte, l)
	for i := 0; i < l; i++ {
		bytes[i] = byte(RandInt(65, 90))
	}
	return string(bytes)
}
func RandInt(min int, max int) int {
	return min + rand.Intn(max-min)
}

func Connect(id uuid.UUID) (res string, err error) {

	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	FailOnError(err, "Failed To connect To RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	FailOnError(err, "Failed To Open a Channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"",    // name
		false, // durable
		false, // delete when unused
		true,  // exclusive
		false, // noWait
		nil,   // arguments
	)
	FailOnError(err, "Failed To Declare a Queue")

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		false,  // auto-ack
		false,  // exclusive
		false,  // no-local
		true,   // no-wait
		nil,    // args
	)
	FailOnError(err, "Failed To Register a Consumer")
	const duration = 3 * time.Second
	timer := time.NewTimer(duration)
	for {
		select {
		case d := <-msgs:
			timer.Reset(duration)
			fmt.Printf("Received a message: %s\n", d.Body)
		case <-timer.C:
			os.Exit(1)
		}
	}

	corrID := RandomString(32)

	stringId := id.String()
	err = ch.Publish(
		"",          // exchange
		"rpc_queue", // routing key
		false,       // mandatory
		false,       // immediate
		amqp.Publishing{
			ContentType:   "text/plain",
			CorrelationId: corrID,
			ReplyTo:       q.Name,
			Body:          []byte(stringId),
		})
	FailOnError(err, "Failed To Publish The Ticket")

	for d := range msgs {
		log.Printf(" [x] %s", d.Body)
		FailOnError(err, "Failed To Send The Ticket")
		break
	}
	return
}
