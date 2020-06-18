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

// use n for loading too, to limit the number of cards loaded from the api
// handle 403 errors in case we are making too many requests somehow
// change load to "prefetch", change load flag so that it simply adds words
// to the local db as pending words
// change export feature so that it if there are no ready to export cards it will try to
// prefetch pending words
// effectively, you would rarely want to run "prefetch" on its own, and prefer adding words
// in pending state and let export to run "prefetch" for you
// but, in case you want to prefetch your words

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
