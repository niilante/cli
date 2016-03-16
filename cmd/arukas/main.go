package main

import (
	arukas "github.com/arukasio/cli"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	arukas.Run(os.Args)
}
