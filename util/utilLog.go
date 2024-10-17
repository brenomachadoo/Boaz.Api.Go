package util

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	exceptionsUtil "bmachado/Boaz.Api.Go/exceptions/generic"

	"github.com/joho/godotenv"
)

func LogInfoMsg(msg string) string {
	return msg
}

func LogInfo(obj interface{}) {
	CheckEnvFileExist()
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
	CheckEnvFileExist()
	var logActiveText = os.Getenv("LOG_SERVICE")
	logActive, err := strconv.ParseBool(logActiveText)
	if err != nil {
		log.Fatal(err)
	}
	if logActive && obj != nil {
		log.Fatal(obj)
	}
}

func CheckEnvFileExist() (bool, error) {

	curDir, osErr := os.Getwd()
	if osErr != nil {
		return false, exceptionsUtil.New("Error loading .env file: " + osErr.Error())
	}

	var newCurDir = strings.Replace(curDir, "\\cmd\\boazApi", "", 1)

	err := godotenv.Load(newCurDir + "\\.env")
	if err != nil {
		return false, exceptionsUtil.New("Error loading .env file: " + err.Error())
	}

	return true, nil
}
