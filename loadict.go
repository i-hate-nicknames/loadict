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
	response, err := fetchWord(appId, appKey, "entail")
	if err != nil {
		log.Fatal(err)
	}
	// marshaled, _ := json.MarshalIndent(response, "", "    ")
	// fmt.Println(string(marshaled))
	text, err := renderCard(response)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(text)
}
