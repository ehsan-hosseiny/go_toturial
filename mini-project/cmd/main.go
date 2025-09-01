package main

import (
	"fmt"
	"log"
	"mini/config"
	"mini/database"
	"mini/server"
)

func main() {
	err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Error loading ,env file %v", err)
	}

	conn, err := database.Connect(config.AppConfig)
	defer database.Close(conn)

	if err != nil {
		log.Fatal(err)
	}

	_, version := database.ExampleQuery(conn)
	fmt.Println(version)

	err = server.StartServer(config.AppConfig, conn)
	if err != nil {
		log.Fatal(err)
	}

}
