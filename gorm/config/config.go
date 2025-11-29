package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	AppName  string
	AppPort  string
	Database database
}

type database struct {
	PostgresUser     string
	PostgresPassword string
	PostgresHost     string
	PostgresDb       string
	PostgresPort     string
}

var AppConfig Config

func LoadConfig() error {
	var cfg Config

	err := godotenv.Load(".env")
	if err != nil {
		return err
	}

	cfg.AppPort = os.Getenv("APP_PORT")
	cfg.AppName = os.Getenv("APP_NAME")
	cfg.Database.PostgresDb = os.Getenv("POSTGRES_DB")
	cfg.Database.PostgresHost = os.Getenv("POSTGRES_HOST")
	cfg.Database.PostgresPassword = os.Getenv("POSTGRES_PASSWORD")
	cfg.Database.PostgresPort = os.Getenv("POSTGRES_PORT")
	cfg.Database.PostgresUser = os.Getenv("POSTGRES_USER")
	AppConfig = cfg
	return nil

}
