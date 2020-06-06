package main

import (
	"encoding/csv"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	"nvm.ga/loadict/fetch"
)

func generateCards(words []string, db *gorm.DB, fetcher fetch.WordFetcher) {

	// todo: check which of the words are in the db, take them from there
	// update last fetched field

	// todo: filter out those words that we have from words slice and
	// fetch only those that we do not have yet

	rendered := fetch.FetchCards(words, fetcher)

	// todo: merge cards together and save them to the db

	// todo: dump merged cards to the file

	file, err := os.Create(outFileName)
	if err != nil {
		log.Fatal("Cannot create file")
	}

	writer := csv.NewWriter(file)

	for card := range rendered {
		err := writer.Write([]string{card.Word, card.Back})
		if err != nil {
			log.Println(err)
		}
	}
	writer.Flush()
}
