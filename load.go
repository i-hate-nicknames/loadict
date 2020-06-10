package main

import (
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	"nvm.ga/loadict/fetch"
)

// loadWords takes list of words from stdin, each word on its own line,
// loads definitions of these words using dictionary API, generates
// html card body using response data and saves these cards to the db
func loadWords(db *gorm.DB) {

	fmt.Println("loading words")
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	appID, appKey := os.Getenv("APP_ID"), os.Getenv("APP_KEY")
	if appID == "" || appKey == "" {
		log.Fatal("Provide app id and app key in .env file")
	}

	// todo: read from stdin
	words := []string{"object", "curtail"}
	fetcher := fetch.MakeFetcher(appID, appKey)

	// todo: check which of the words are in the db, take them from there
	// update last fetched field

	// todo: filter out those words that we have from words slice and
	// fetch only those that we do not have yet

	cards := fetch.FetchCards(words, fetcher)

	fmt.Printf("Fetched %d cards!\n", len(cards))
	// todo: merge cards together and save them to the db

}
