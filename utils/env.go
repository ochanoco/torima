package utils

import (
	"fmt"
	"log"
	"os"
)

func ReadEnv(name, def string) string {
	value := os.Getenv(name)

	if value == "" {
		fmt.Printf("environment variable '%v' is not found so that proxy use '%v'\n", name, def)
		value = def
	}

	return value
}

func ReadEnvOrPanic(name string) string {
	value := os.Getenv(name)

	if value == "" {
		log.Fatalf("environment variable '%v' is not found", name)
	}

	return value
}
