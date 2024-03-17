package env

import (
	"fmt"
	"github.com/joho/godotenv"
	"network-challenges/src/utils/log"
	"os"
	"strconv"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Error(fmt.Errorf("no .env file found"))
	}
}

func get(key string) (string, error) {
	value, exists := os.LookupEnv(key)

	if exists {
		return value, nil
	} else {
		return "", fmt.Errorf("key %s not found", key)
	}
}

func GetSmokeTestPort() int {
	port, err := get("SMOKE_TEST_PORT")
	if err != nil {
		return 3000
	}
	value, err := strconv.Atoi(port)
	if err != nil {
		return 3000
	}
	return value
}

func GetPrimeTimePort() int {
	port, err := get("PRIME_TIME_PORT")
	if err != nil {
		return 4000
	}
	value, err := strconv.Atoi(port)
	if err != nil {
		return 4000
	}
	return value
}

func GetMeansEndPort() int {
	port, err := get("MEANS_END_PORT")
	if err != nil {
		return 5000
	}
	value, err := strconv.Atoi(port)
	if err != nil {
		return 5000
	}
	return value
}
