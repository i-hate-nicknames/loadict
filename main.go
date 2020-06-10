package main

import (
	"flag"
	"log"

	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"nvm.ga/loadict/db"
)

const exportFile = "cards.csv"

const exportCardsNum = 10

var load = flag.Bool("load", false, "loads list of words and stores them locally")
var export = flag.Bool("export", false, "exports a number of card into csv file")
var exportNumber = flag.Int("n", exportCardsNum, "number of cards to export, default 10")

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
		exportCards(*exportNumber, conn)
	}
}
