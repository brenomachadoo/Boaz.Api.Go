package util

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

func LogInfo(obj interface{}) {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Print("Error loading .env file: " + err.Error())
	}
	var logActiveText = os.Getenv("LOG_SERVICE")
	logActive, err := strconv.ParseBool(logActiveText)
	if err != nil {
		log.Fatal(err)
	}
	if logActive {
		fmt.Print(obj)
	}
}

func LogFatal(obj interface{}) {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Print("Error loading .env file: " + err.Error())
	}
	var logActiveText = os.Getenv("LOG_SERVICE")
	logActive, err := strconv.ParseBool(logActiveText)
	if err != nil {
		log.Fatal(err)
	}
	if logActive && obj != nil {
		log.Fatal(obj)
	}
}
