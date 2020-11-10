package consume

import (
	"log"
	"task/logic/ticketFunctions"
	"task/model"

	uuid "github.com/satori/go.uuid"
	"github.com/streadway/amqp"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

//func BytesToString(b []byte) string {
//	bh := (*reflect.SliceHeader)(unsafe.Pointer(&b))
//	sh := reflect.StringHeader{bh.Data, bh.Len}
//	return *(*string)(unsafe.Pointer(&sh))
//}

//Consume tyo listen to the queue
func Consume() {

	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "Failed To Connect To RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed To Open a Channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"rpc_queue", // name
		false,       // durable
		false,       // delete when unused
		false,       // exclusive
		false,       // no-wait
		nil,         // arguments
	)
	failOnError(err, "Failed To Declare a Queue")

	err = ch.Qos(
		1,     // prefetch count
		0,     // prefetch size
		false, // global
	)
	failOnError(err, "Failed to set QoS")

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		false,  // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)

	failOnError(err, "Failed To Register a Consumer")

	forever := make(chan bool)
	go func() {
		for d := range msgs {

			str := string(d.Body)
			var TId = uuid.Must(uuid.FromString(str))
			var msg, _ = ticket.GetTicketById(TId)
			log.Printf(" [x] %s", msg)

			response := model.Processing

			ticket.UpdateOperatorIDAndStatus(TId, ticket.OperatorID, response)

			err = ch.Publish(
				"",        // exchange
				d.ReplyTo, // routing key
				false,     // mandatory
				false,     // immediate
				amqp.Publishing{
					ContentType:   "text/plain",
					CorrelationId: d.CorrelationId,
					Body:          []byte(response),
				})
			failOnError(err, "Failed To Publish The Response")
			d.Ack(false)
		}
	}()

	log.Printf(" [*] Awaiting Ticket Requests")
	<-forever
}
