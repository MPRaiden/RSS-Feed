package main

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main () {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Coul not load env file")
	}

	jwtSecret := os.Getenv("PORT")
	if jwtSecret == "" {
		log.Fatal("Unable to load PORT env variable.")
	}
}
