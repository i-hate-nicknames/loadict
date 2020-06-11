package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	"nvm.ga/loadict/fetch"
)

// loadWords takes list of words from stdin, each word on its own line,
// loads definitions of these words using dictionary API, generates
// html card body using response data and saves these cards to the db
func loadWords(conn *gorm.DB) {

	fmt.Println("loading words")
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	appID, appKey := os.Getenv("APP_ID"), os.Getenv("APP_KEY")
	if appID == "" || appKey == "" {
		log.Fatal("Provide app id and app key in .env file")
	}

	words := readWords()
	fetcher := fetch.MakeFetcher(appID, appKey)

	// todo: check which of the words are in the db, take them from there
	// update last fetched field

	// todo: filter out those words that we have from words slice and
	// fetch only those that we do not have yet

	log.Println("Loading the following words:", words)
	cards := fetch.FetchCards(words, fetcher)

	// cards := make([]*card.Card, 0)
	// cards = append(cards, card.MakeCard("test", "test_back"))
	// cards = append(cards, card.MakeCard("test2", "test_back2"))

	fmt.Printf("Fetched %d cards!\n", len(cards))

	// err = db.SaveCards(conn, cards)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// todo: merge cards together and save them to the db

}

func readWords() []string {
	reader := bufio.NewReader(os.Stdin)
	words := make([]string, 0)
	for {
		word, err := reader.ReadString('\n')
		if err == io.EOF {
			return words
		}
		if err != nil {
			log.Fatalln("Cannot read words:", err.Error())
		}
		words = append(words, strings.TrimSpace(word))
	}
}
