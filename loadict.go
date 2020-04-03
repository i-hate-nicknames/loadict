package main

import (
	"encoding/csv"
	"log"
	"os"

	"github.com/joho/godotenv"
)

// todo: read comma-separated list of words
var word = "entail"

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	appId, appKey := os.Getenv("APP_ID"), os.Getenv("APP_KEY")
	if appId == "" || appKey == "" {
		log.Fatal("Provide app id and app key in .env file")
	}
	response, err := fetchWord(appId, appKey, word)
	if err != nil {
		log.Fatal(err)
	}
	// marshaled, _ := json.MarshalIndent(response, "", "    ")
	// fmt.Println(string(marshaled))
	text, err := renderCard(response)
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Println(text)
	export(word, text)
}

func export(word, value string) error {
	writer := csv.NewWriter(os.Stdout)
	err := writer.Write([]string{word, value})
	if err != nil {
		return err
	}
	writer.Flush()
	return nil
}
