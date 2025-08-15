package config

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type Config struct {
	PostgresqlUser     string
	PostgresqlPassword string
	PostgresqlHost     string
	PostgresqlDb       string
}

var GlobalConfig Config

func LoadEnv() {
	file, err := os.Open(".env")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	config := Config{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.TrimSpace(line) == "" {
			continue
		}
		parts := strings.SplitN(line, "=", 2)
		if len(parts) != 2 {
			continue
		}
		key := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(parts[1])
		fmt.Println(key, " - ", value)

		switch key {
		case "POSTGRESQL_USERNAME":
			config.PostgresqlUser = value
		case "POSTGRESQL_PASSWORD":
			config.PostgresqlPassword = value
		case "POSTGRESQL_HOST":
			config.PostgresqlHost = value
		case "POSTGRESQL_DB":
			config.PostgresqlDb = value
		}
	}

	if scanner.Err() != nil {
		log.Fatal(scanner.Err())
	}

	GlobalConfig = config

}
