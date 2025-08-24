package main

import (
	"fmt"
	"log"
	"mini/config"
	"mini/database"
)

func main() {
	err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Error loading ,env file %v", err)
	}

	conn, err := database.Connect(config.AppConfig)

	if err != nil {
		log.Fatal(err)
	}

	err, version := database.ExampleQuery(conn)

	fmt.Println(version)

}
