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
	appId, appKey := os.Getenv("APP_ID"), os.Getenv("APP_KEY")
	if appId == "" || appKey == "" {
		log.Fatal("Provide app id and app key in .env file")
	}
	fmt.Println("App id: " + appId)
	fmt.Println("Words loader for anki will be here some day")
}
