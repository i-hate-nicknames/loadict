package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"nvm.ga/loadict/db"
	"nvm.ga/loadict/fetch"
)

const outFileName = "cards.csv"

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	appID, appKey := os.Getenv("APP_ID"), os.Getenv("APP_KEY")
	if appID == "" || appKey == "" {
		log.Fatal("Provide app id and app key in .env file")
	}

	conn := db.Connect()
	db.Migrate(conn)

	// todo: read from arguments
	words := []string{"object"}
	fetcher := fetch.MakeFetcher(appID, appKey)
	generateCards(words, conn, fetcher)
}
