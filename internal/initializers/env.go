package initializers

import (
	"github.com/joho/godotenv"
	"log"
)

func LoadConfig() error {
	err := godotenv.Load("../../.env")
	if err != nil {
	   log.Println("Failed to load config from env because", err.Error())
	}
	return nil 
}