package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	appId := os.Getenv("APP_ID")
	fmt.Println("App id: " + appId)
	fmt.Println("Words loader for anki will be here some day")
}
