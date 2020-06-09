package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/joho/godotenv"
	"nvm.ga/loadict/db"
	"nvm.ga/loadict/fetch"
)

const exportFile = "cards.csv"

const exportCards = 10

var load = flag.Bool("load", false, "loads list of words and stores them locally")
var export = flag.Bool("export", false, "exports a number of card into csv file")
var exportNumber = flag.Int("n", exportCards, "number of cards to export, default 10")

func main() {
	flag.Parse()
	if !*load && !*export {
		log.Panicln("Please, either load or export cards")
	}
	if *load && *export {
		log.Panicln("Please, either load or export cards")
	}
	conn := db.Connect()
	db.Migrate(conn)
	if *load {
		loadWords(conn)
	} else {
		exportWords(*exportNumber, conn)
	}
}

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

	// todo: read from stdin or smth
	words := []string{"object"}
	fetcher := fetch.MakeFetcher(appID, appKey)

	// todo: take out export functionality to export function
	generateCards(words, conn, fetcher)
}

func exportWords(n int, conn *gorm.DB) {
	fmt.Printf("Exporting %d words to %s\n", n, exportFile)
}
