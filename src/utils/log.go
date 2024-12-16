package utils

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

func getEnvVar(key string) string {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal(err)
	}

	return os.Getenv(key)
}

// Debug logs a message if the DEBUG_LEVEL is greater than 0
func Debug(message ...any) {
	debug, err := strconv.Atoi(getEnvVar("DEBUG_LEVEL"))
	if err != nil {
		debug = 0
	}

	if debug > 0 {
		args := []any{"[DEBUG]"}
		args = append(args, message...)
		log.Println(args...)
	}
}
