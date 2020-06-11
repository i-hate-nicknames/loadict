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
	"nvm.ga/loadict/card"
	"nvm.ga/loadict/db"
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
	existingCards := db.LoadCardsByWords(conn, words)

	wordsToFetch := filterExisting(words, existingCards)

	if len(wordsToFetch) == 0 {
		log.Fatal("All the words are already loaded")
	}

	log.Println("Loading the following words:", wordsToFetch)
	cards := fetch.FetchCards(wordsToFetch, fetcher)

	fmt.Printf("Fetched %d cards!\n", len(cards))

	err = db.SaveCards(conn, cards)
	if err != nil {
		log.Fatal(err)
	}
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

func filterExisting(words []string, existing []*card.Card) []string {
	result := make([]string, 0)
	existingMap := make(map[string]bool, 0)
	for _, card := range existing {
		existingMap[card.Word] = true
	}
	for _, word := range words {
		if !existingMap[word] {
			result = append(result, word)
		}
	}
	return result
}
