package main

import (
	"fmt"

	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {

	conn, err := amqp.Dial("amqp://admin:123456@localhost:5673")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		fmt.Println(err)
		return
	}

	defer ch.Close()

	queue, err := ch.QueueDeclare("test", false, false, false, false, nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	for i := 0; i < 100; i++ {
		err = ch.Publish("", queue.Name, false, false, amqp.Publishing{

			ContentType: "text/plain",
			Body:        []byte(fmt.Sprintf("msg-%d", i)),
		})
		if err != nil {
			fmt.Println(err)
			return
		}

	}

	msgs, err := ch.Consume(queue.Name, "", true, false, false, false, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	for d := range msgs {
		fmt.Println(string(d.Body))
	}

}
