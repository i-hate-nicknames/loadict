package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	"nvm.ga/loadict/db"
)

// Export at most num cards from the db into .csv file, ready
// to be imported by anki
// Exported cards will be marked as such in the db and won't be
// exported in future
func exportCards(num int, conn *gorm.DB) {
	fmt.Printf("Exporting %d words to %s\n", num, exportFile)
	file, err := os.Create(exportFile)
	if err != nil {
		log.Fatal("Cannot create file")
	}

	writer := csv.NewWriter(file)

	cards := db.LoadCards(conn, num)

	for _, card := range cards {
		err := writer.Write([]string{card.Word, card.Back})
		if err != nil {
			log.Println(err)
		}
		card.Exported = true
	}
	writer.Flush()

	// todo: save cards to the db to reflect that they have been exported
}
