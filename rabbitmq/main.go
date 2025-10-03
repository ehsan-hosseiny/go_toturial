package main

import (
	"fmt"
	"log"
	"rabbitmq/config"
	"rabbitmq/rabbitmq"
)

func main() {

	err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}

	r, err := rabbitmq.Connect(config.AppConfig)
	if err != nil {
		log.Fatal(err)
	}
	q, err := r.DeclareQueue("test")
	if err != nil {
		log.Fatal(err)
	}
	err = r.PublishMessage(q.Name, "hello from rabbitmq")
	if err != nil {
		log.Fatal(err)
	}

	messages, err := r.Consumer(q.Name)
	if err != nil {
		log.Fatal(err)
	}
	for d := range messages {
		fmt.Println(string(d.Body))
	}

}
