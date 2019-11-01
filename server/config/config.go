package config

import (
	"os"

	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
)

type Config struct {
	AppPort        string
	ElasticAddress string
	BookingIndex   string
}

var GlobalConfig Config

func init() {
	if err := godotenv.Load(); err != nil {
		log.Warn("Error loading .env file")
	}
	GlobalConfig = Config{
		AppPort:        os.Getenv("APP_PORT"),
		ElasticAddress: os.Getenv("ELASTIC_ADDRESS"),
		BookingIndex:   os.Getenv("BOOKING_INDEX"),
	}
}
